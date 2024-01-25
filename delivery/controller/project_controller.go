package controller

import (
	"log"
	"net/http"
	"strconv"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/common"
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	projectUsecase usecase.ProjectUseCase
	rg             *gin.RouterGroup
}

func NewProjectController(projectUsecase usecase.ProjectUseCase, rg *gin.RouterGroup) *ProjectController {
	return &ProjectController{
		projectUsecase: projectUsecase,
		rg:             rg,
	}
}

func (a *ProjectController) Route() {
	a.rg.GET("/project", a.GetAll)
	a.rg.GET("/project/id/:id", a.GetProjectById)
	a.rg.GET("/project/deadline/:deadline", a.GetProjectsByDeadline)
	a.rg.GET("/project/manager/:id", a.GetProjectsByManagerId)
	a.rg.GET("/project/member/:id", a.GetProjectsByMemberId)
	a.rg.POST("/project/create", a.CreateNewProject)
	a.rg.POST("/project/addmember/:id", a.AddProjectMember)
	a.rg.DELETE("/project/deletemember/:id", a.DeleteProjectMember)
	a.rg.GET("/project/allmember/:id", a.GetAllProjectMember)
	a.rg.PUT("/project/update", a.UpdateProject)
	a.rg.DELETE("/project/delete/:id", a.DeleteProject)

}

func (pc *ProjectController) GetAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid page parameter")
		return
	}

	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid size parameter")
		return
	}

	projects, paging, err := pc.projectUsecase.GetAll(page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully retrieved projects")

	response := map[string]interface{}{
		"projects": projects,
		"paging":   paging,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Get Resource",
		"data":    response,
	})
}

func (pc *ProjectController) GetProjectById(c *gin.Context) {
	id := c.Param("id")
	project, err := pc.projectUsecase.GetProjectById(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Project id"+id+"not found")
		return
	}

	log.Printf("succes Get Resource")

	response := map[string]interface{}{
		"project": project,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Get Resource",
		"data":    response,
	})
}

func (pc *ProjectController) GetProjectsByDeadline(c *gin.Context) {
	// Parse deadline parameter from query string
	deadline := c.Param("deadline")

	projects, err := pc.projectUsecase.GetProjectsByDeadline(deadline)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Error get project by Deadline")
		return
	}

	log.Printf("Successfully retrieved projects by deadline: %s", deadline)

	response := map[string]interface{}{
		"projects": projects,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Get Resource",
		"data":    response,
	})
}

func (pc *ProjectController) GetProjectsByManagerId(c *gin.Context) {
	managerID := c.Param("id")

	projects, err := pc.projectUsecase.GetProjectsByManagerId(managerID)
	if err != nil {
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
		common.SendErrorResponse(c, http.StatusInternalServerError, "Error get Project by Member Id")
		return
	}

	log.Printf("Successfully retrieved projects by member ID: %s", memberID)

	common.SendSingleResponse(c, projects, "Success")
}

func (pc *ProjectController) CreateNewProject(c *gin.Context) {
	var request model.Project

	if err := c.ShouldBindJSON(&request); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	createdProject, err := pc.projectUsecase.CreateNewProject(request)
	if err != nil {
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
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := pc.projectUsecase.AddProjectMember(id, request.Members)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully added project members to project ID: %s", id)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Add Project Members",
	})
}

func (pc *ProjectController) DeleteProjectMember(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		Members []string `json:"members"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := pc.projectUsecase.DeleteProjectMember(id, request.Members)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully deleted project members from project ID: %s", id)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Delete Project Members",
	})
}

func (pc *ProjectController) GetAllProjectMember(c *gin.Context) {
	id := c.Param("id")

	members, err := pc.projectUsecase.GetAllProjectMember(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully retrieved all project members for project ID: %s", id)

	response := map[string]interface{}{
		"members": members,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Get All Project Members",
		"data":    response,
	})
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {
	var request model.Project

	if err := c.ShouldBindJSON(&request); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedProject, err := pc.projectUsecase.Update(request)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully updated project with ID: %s", updatedProject.Id)

	response := map[string]interface{}{
		"project": updatedProject,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Update Project",
		"data":    response,
	})
}

func (pc *ProjectController) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	err := pc.projectUsecase.Delete(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Successfully deleted project with ID: %s", id)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success Delete Project",
	})
}
