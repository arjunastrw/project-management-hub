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

type TaskController struct {
	taskUC         usecase.TaskUsecase
	authMiddleware middleware.AuthMiddleware
	rg             *gin.RouterGroup
}

func NewTaskController(taskUC usecase.TaskUsecase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup) *TaskController {
	return &TaskController{
		taskUC:         taskUC,
		authMiddleware: authMiddleware,
		rg:             rg,
	}
}

// update ini
func (t *TaskController) Route() {
	t.rg.GET("/tasks/list", t.authMiddleware.RequireToken("ADMIN", "MANAGER"), t.GetAllTask)
	t.rg.GET("/tasks/getbypic/:id", t.authMiddleware.RequireToken("ADMIN", "MANAGER", "TEAM MEMBER"), t.GetTaskByPersonInCharge)
	t.rg.GET("/tasks/getbyid/:id", t.authMiddleware.RequireToken("ADMIN", "MANAGER", "TEAM MEMBER"), t.GetTaskById)
	t.rg.GET("/tasks/getbyprojectid/:id", t.authMiddleware.RequireToken("ADMIN", "MANAGER", "TEAM MEMBER"), t.GetTaskByProjectId)
	t.rg.POST("/tasks/create", t.authMiddleware.RequireToken("MANAGER"), t.CreateTask)
	t.rg.PUT("/tasks/update/:id", t.authMiddleware.RequireToken("MANAGER", "TEAM MEMBER"), t.UpdateTask)
	t.rg.DELETE("/tasks/delete/:id", t.authMiddleware.RequireToken("MANAGER"), t.DeleteTask)
}

func (t *TaskController) CreateTask(c *gin.Context) {
	var newtask model.Task
	if err := c.ShouldBind(&newtask); err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := t.taskUC.CreateTask(newtask)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("Success: ")
	common.SendSingleResponse(c, task, "Success")
}

func (t *TaskController) GetTaskByPersonInCharge(c *gin.Context) {
	pic := c.Param("id")
	tasks, err := t.taskUC.GetByPersonInCharge(pic)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, "tasks by pic_id "+pic+" not found")
		return
	}

	log.Println("Success: ")
	common.SendSingleResponse(c, tasks, "Success")

}

func (t *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := t.taskUC.GetById(id)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, "tasks id "+id+" not found")
		return
	}

	log.Println("Success: ")
	common.SendSingleResponse(c, task, "Success")

}

func (t *TaskController) GetTaskByProjectId(c *gin.Context) {
	projectid := c.Param("id")
	tasks, err := t.taskUC.GetByProjectId(projectid)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, "tasks with project id "+projectid+" not found")
		return
	}

	log.Println("Success: ")
	common.SendSingleResponse(c, tasks, "Success")

}

func (t *TaskController) GetAllTask(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	tasks, paging, err := t.taskUC.GetAll(page, size)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, "no task found")
		return
	}

	var resp []interface{}

	for _, v := range tasks {
		resp = append(resp, v)
	}

	log.Println("Success: ")
	common.SendPagedResponse(c, resp, paging, "OK")
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	userId := c.Param("id")
	var newtask model.Task
	if err := c.ShouldBind(&newtask); err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//disini cekrole if manager >> updattaskbymanager, if pic >> updatetaskbymember
	task, err := t.taskUC.UpdateTask(userId, newtask)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("Success: ")
	common.SendSingleResponse(c, task, "Success")
}

func (t *TaskController) DeleteTask(c *gin.Context) {

	id := c.Param("id")

	err := t.taskUC.Delete(id)
	if err != nil {
		log.Println(err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("Success: ")
	common.SendSingleResponse(c, nil, "Success")
}
