package controller

import (
	"enigma.com/projectmanagementhub/delivery/middleware"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/common"
	"enigma.com/projectmanagementhub/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type UserController struct {
	userUC         usecase.UserUseCase
	authMiddleware middleware.AuthMiddleware
	rg             *gin.RouterGroup
}

func NewUserController(rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware, userUC usecase.UserUseCase) *UserController {
	return &UserController{
		userUC:         userUC,
		authMiddleware: authMiddleware,
		rg:             rg,
	}
}

func (a *UserController) Route() {
	a.rg.GET("/user/list", a.authMiddleware.RequireToken("ADMIN"), a.FindAllUser)
	a.rg.GET("/user/:id", a.authMiddleware.RequireToken("ADMIN", "MANAGER", "TEAM MEMBER"), a.FindUserById)
	a.rg.GET("/user/email/:email", a.authMiddleware.RequireToken("ADMIN", "MANAGER", "TEAM MEMBER"), a.FindUserByEmail)
	a.rg.POST("/user/create", a.authMiddleware.RequireToken("ADMIN"), a.CreateUser)
	a.rg.PUT("/user/update", a.authMiddleware.RequireToken("ADMIN"), a.UpdateUser)
	a.rg.DELETE("/user/delete/:id", a.authMiddleware.RequireToken("ADMIN"), a.DeleteUser)
}

func (a *UserController) FindAllUser(c *gin.Context) {
	// Get parameters from query
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	// Call FindAllUser method
	users, paging, err := a.userUC.FindAllUser(page, size)
	if err != nil {
		//log bad request
		log.Println("Failed to get users" + err.Error())
		// return bad request
		common.SendErrorResponse(c, 400, "Failed to get users")
		return
	}
	var resp []interface{}

	for _, v := range users {
		resp = append(resp, v)
	}
	// log success
	log.Println("Success Get Resource")
	// return success
	common.SendPagedResponse(c, resp, paging, "OK")
}

func (a *UserController) FindUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := a.userUC.FindUserById(id)
	if err != nil {
		// Log For Error
		log.Println("Failed to get users" + err.Error())
		// Return Bad Request
		common.SendErrorResponse(c, 400, "User with ID "+id+" not found")
		return
	}
	// Log if success
	log.Println("Success Get Resource")
	// Return Success
	common.SendSingleResponse(c, user, "Success")
}

func (a *UserController) FindUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := a.userUC.FindUserByEmail(email)
	if err != nil {
		// Log For Error
		log.Println("Failed to get users" + err.Error())
		// Return Bad Request
		common.SendErrorResponse(c, 400, "User With Email "+email+" Not Found")
		return
	}

	// Log if success
	log.Println("Success Get Resource")
	// Return Success
	common.SendSingleResponse(c, user, "Success")
}

func (a *UserController) CreateUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBind(&newUser); err != nil {
		// Log For Bad Request
		log.Println("Failed to bind JSON: " + err.Error())
		// Return Bad Request
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	user, err := a.userUC.CreateUser(newUser)
	if err != nil {
		// Log For Error
		log.Println("Failed to create user: " + err.Error())
		// Return Internal Server Error
		common.SendErrorResponse(c, 500, err.Error())
		return
	}
	// Log For Success
	log.Println("Success Create User")
	// Return Success
	common.SendSingleResponse(c, user, "Success")
}

func (a *UserController) UpdateUser(c *gin.Context) {

	// Bind JSON request ke model User
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		// Log For Bad Request
		log.Println("Failed to bind JSON: " + err.Error())
		// Return Bad Request
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	// Update User
	updatedUser, err := a.userUC.UpdateUser(user)
	if err != nil {
		// Log For Update User Error
		log.Println("Failed to update user: " + err.Error())
		// Return Internal Server Error
		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	// Log For Success
	log.Println("Success Update User")
	// Return Success
	common.SendSingleResponse(c, updatedUser, "Success")
}

func (a *UserController) DeleteUser(c *gin.Context) {
	// Get ID from URL parameter
	id := c.Param("id")

	// Check if ID is empty
	if id == "" {
		// Log For Error
		log.Println("Error occurred")
		// Return Bad Request
		common.SendErrorResponse(c, 400, "ID is required")
		return
	}

	// Delete User
	err := a.userUC.DeleteUser(id)
	if err != nil {
		// Log For Error
		log.Println("Failed to delete user: " + err.Error())
		// Return Internal Server Error
		common.SendErrorResponse(c, 500, fmt.Sprintf("Failed to delete user: %s", err))
		return
	}

	// log for success
	log.Println("Success Delete User")
	// return success
	common.SendSingleResponse(c, nil, "Success Delete User")
}
