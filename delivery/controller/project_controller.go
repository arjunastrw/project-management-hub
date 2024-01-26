package controller

import (
	"log"
	"net/http"
	"strconv"

	"enigma.com/projectmanagementhub/delivery/middleware"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/common"
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	projectUsecase usecase.ProjectUseCase
	authMiddleware middleware.AuthMiddleware
	rg             *gin.RouterGroup
}

func NewProjectController(projectUsecase usecase.ProjectUseCase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup) *ProjectController {
	return &ProjectController{
		projectUsecase: projectUsecase,
		authMiddleware: authMiddleware,
		rg:             rg,
	}
}

func (a *ProjectController) Route() {
	a.rg.GET("/project", a.authMiddleware.RequireToken("ADMIN"), a.GetAll)
	a.rg.GET("/project/id/:id", a.authMiddleware.RequireToken("ADMIN", "MANAGER", "TEAM MEMBER"), a.GetProjectById)
	a.rg.GET("/project/deadline/:deadline", a.authMiddleware.RequireToken("ADMIN", "MANAGER"), a.GetProjectsByDeadline)
	a.rg.GET("/project/manager/:id", a.authMiddleware.RequireToken("ADMIN", "MANAGER"), a.GetProjectsByManagerId)
	a.rg.GET("/project/member/:id", a.authMiddleware.RequireToken("ADMIN", "MANAGER", "TEAM MEMBER"), a.GetProjectsByMemberId)
	a.rg.POST("/project/create", a.authMiddleware.RequireToken("ADMIN"), a.CreateNewProject)
	a.rg.POST("/project/addmember/:id", a.authMiddleware.RequireToken("ADMIN", "MANAGER"), a.AddProjectMember)
	a.rg.DELETE("/project/deletemember/:id", a.authMiddleware.RequireToken("ADMIN", "MANAGER"), a.DeleteProjectMember)
	a.rg.GET("/project/allmember/:id", a.authMiddleware.RequireToken("ADMIN", "MANAGER", "TEAM MEMBER"), a.GetAllProjectMember)
	a.rg.PUT("/project/update", a.authMiddleware.RequireToken("ADMIN", "MANAGER"), a.UpdateProject)
	a.rg.DELETE("/project/delete/:id", a.authMiddleware.RequireToken("ADMIN"), a.DeleteProject)

}

func (pc *ProjectController) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid page parameter")
		return
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid size parameter")
		return
	}

	projects, paging, err := pc.projectUsecase.GetAll(page, size)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully retrieved projects")

	var projectsInterfaceSlice []interface{}
	for _, project := range projects {
		projectsInterfaceSlice = append(projectsInterfaceSlice, project)
	}

	common.SendPagedResponse(c, projectsInterfaceSlice, paging, "Success Get Resource")
}

func (pc *ProjectController) GetProjectById(c *gin.Context) {
	id := c.Param("id")
	project, err := pc.projectUsecase.GetProjectById(id)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, "Project id"+id+"not found")
		return
	}

	log.Printf("succes Get Resource")

	common.SendSingleResponse(c, map[string]interface{}{
		"project": project,
	}, "Success Get Resource")
}

func (pc *ProjectController) GetProjectsByDeadline(c *gin.Context) {
	// Parse deadline parameter from query string
	deadline := c.Param("deadline")

	projects, err := pc.projectUsecase.GetProjectsByDeadline(deadline)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, "Error get project by Deadline")
		return
	}

	log.Printf("Successfully retrieved projects by deadline: %s", deadline)

	common.SendSingleResponse(c, projects, "Success Get Resource")
}

func (pc *ProjectController) GetProjectsByManagerId(c *gin.Context) {
	managerID := c.Param("id")

	projects, err := pc.projectUsecase.GetProjectsByManagerId(managerID)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, "Error get Project Manager by Id")
		return
	}

	log.Printf("Successfully retrieved projects by manager ID: %s", managerID)

	common.SendSingleResponse(c, projects, "Success")
}

func (pc *ProjectController) GetProjectsByMemberId(c *gin.Context) {
	memberID := c.Param("id")

	projects, err := pc.projectUsecase.GetProjectsByMemberId(memberID)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, "Error get Project by Member Id")
		return
	}

	log.Printf("Successfully retrieved projects by member ID: %s", memberID)

	common.SendSingleResponse(c, projects, "Success")
}

func (pc *ProjectController) CreateNewProject(c *gin.Context) {
	var request model.Project

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	createdProject, err := pc.projectUsecase.CreateNewProject(request)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully created new project with ID: %s", createdProject.Id)

	common.SendSingleResponse(c, createdProject, "Success")
}

func (pc *ProjectController) AddProjectMember(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		Members []string `json:"members"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := pc.projectUsecase.AddProjectMember(id, request.Members)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully added project members to project ID: %s", id)

	common.SendSingleResponse(c, nil, "Success")
}

func (pc *ProjectController) DeleteProjectMember(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		Members []string `json:"members"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := pc.projectUsecase.DeleteProjectMember(id, request.Members)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully deleted project members from project ID: %s", id)

	common.SendSingleResponse(c, nil, "Success")
}

func (pc *ProjectController) GetAllProjectMember(c *gin.Context) {
	id := c.Param("id")

	members, err := pc.projectUsecase.GetAllProjectMember(id)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully retrieved all project members for project ID: %s", id)

	common.SendSingleResponse(c, members, "Success Get Resource")
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {
	var request model.Project

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedProject, err := pc.projectUsecase.Update(request)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully updated project with ID: %s", updatedProject.Id)

	common.SendSingleResponse(c, updatedProject, "Success Get Resource")
}

func (pc *ProjectController) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	err := pc.projectUsecase.Delete(id)
	if err != nil {
		log.Println(err.Error())

		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully deleted project with ID: %s", id)

	common.SendSingleResponse(c, nil, "Success")
}
