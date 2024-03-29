package delivery

import (
	"database/sql"
	"fmt"

	"enigma.com/projectmanagementhub/config"
	"enigma.com/projectmanagementhub/delivery/controller"
	"enigma.com/projectmanagementhub/delivery/middleware"
	"enigma.com/projectmanagementhub/report"
	"enigma.com/projectmanagementhub/shared/service"

	"enigma.com/projectmanagementhub/repository"
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type Server struct {
	userUC     usecase.UserUseCase
	taskUC     usecase.TaskUsecase
	projectUC  usecase.ProjectUseCase
	reportUC   usecase.ReportUsecase
	authUC     usecase.AuthUsecase
	engine     *gin.Engine
	jwtService service.JwtService
	host       string
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/pmh-api/v1")

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewUserController(rg, authMiddleware, s.userUC).Route()
	controller.NewTaskController(s.taskUC, authMiddleware, rg).Route()
	controller.NewProjectController(s.projectUC, authMiddleware, rg).Route()
	controller.NewReportController(s.reportUC, authMiddleware, rg).Route()
	controller.NewAuthController(s.authUC, rg).Route()

}

func NewServer() *Server {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)

	db, err := sql.Open(cfg.Driver, psqlInfo)
	if err != nil {
		panic(err)
	}

	report := report.NewReportToTXT(cfg.PathConfig)

	//inject db ke repository
	taskRepository := repository.NewTaskRepository(db)
	userRepository := repository.NewUserRepository(db)
	projectRepository := repository.NewProjectRepository(db)
	reportRepository := repository.NewReportRepository(db, report)

	//inject repository ke usecase
	UserUseCase := usecase.NewUserUseCase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, userRepository, projectRepository)
	projectUsecase := usecase.NewProjectUseCase(projectRepository, userRepository)
	reportUsecase := usecase.NewReportUsecase(reportRepository, taskRepository)

	jwtService := service.NewJwtService(cfg.TokenConfig)
	authUsecase := usecase.NewAuthUsecase(UserUseCase, jwtService)

	engine := gin.Default()
	host := cfg.ApiPort

	return &Server{
		userUC:     UserUseCase,
		taskUC:     taskUsecase,
		projectUC:  projectUsecase,
		reportUC:   reportUsecase,
		engine:     engine,
		host:       host,
		authUC:     authUsecase,
		jwtService: jwtService,
	}
}
