package controller

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/common"
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUC usecase.UserUseCase
	rg     *gin.RouterGroup
}

func NewUserController(rg *gin.RouterGroup, userUC usecase.UserUseCase) *UserController {
	return &UserController{
		userUC: userUC,
		rg:     rg,
	}
}

func (a *UserController) Route() {
	a.rg.GET("/users/list", a.FindAllUser)
	a.rg.GET("/user/:id", a.FindUserById)
	a.rg.GET("/user/email/:email", a.FindUserByEmail)
	a.rg.POST("/users", a.CreateUser)
	a.rg.PUT("/user/:id", a.UpdateUser)
	a.rg.DELETE("/user/:id", a.DeleteUser)
}

func (a *UserController) FindAllUser(c *gin.Context) {
	// Get parameters from query
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	// Call FindAllUser method
	users, paging, err := a.userUC.FindAllUser(page, size)
	if err != nil {
		common.SendErrorResponse(c, 400, "Failed to get users")
		return
	}

	// Log for Get User Success
	log.Println("Success Get Resource")

	var resp []interface{}

	for _, v := range users {
		resp = append(resp, v)
	}
	common.SendPagedResponse(c, resp, paging, "OK")
}

func (a *UserController) FindUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := a.userUC.FindUserById(id)
	if err != nil {
		// Log For Error

		// Return If Condition Error
		c.JSON(404, gin.H{
			"message": "User with ID " + id + " Not Found" + err.Error(),
		})
		return
	}
	// Validate If User Found

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Success Get Resource",
		"data":    user,
	})
}

func (a *UserController) FindUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := a.userUC.FindUserByEmail(email)
	if err != nil {
		// Log For Error

		// Return If Condition Error
		c.JSON(404, gin.H{
			"message": "User With Email " + email + " Not Found" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Success Get Resource",
		"data":    user,
	})
}

func (a *UserController) CreateUser(c *gin.Context) {
	var newuser model.User
	if err := c.ShouldBind(&newuser); err != nil {
		common.SendErrorResponse(c, 400, err.Error())
		return
	}

	user, err := a.userUC.CreateUser(newuser)
	if err != nil {
		if strings.Contains(err.Error(), "Email already exist") {
			common.SendErrorResponse(c, 400, "Email already exists")
			return
		}

		common.SendErrorResponse(c, 500, err.Error())
		return
	}

	common.SendSingleResponse(c, user, "Success")
}

// func (a *UserController) CreateUser(c *gin.Context) {
// 	// Bind JSON request to User Model
// 	user := model.User{}
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		// Log For Bad Request
// 		logrus.Errorf("Failed to bind JSON: %s", err.Error())
// 		c.JSON(400, gin.H{
// 			"message": "Failed to bind JSON: " + err.Error(),
// 		})
// 		return
// 	}

// 	email := c.Param("email")
// 	// Check If Email Already Exist
// 	existingUser, err := a.userUC.FindUserByEmail(email)
// 	if err != nil {
// 		// Log for Checking Existing User Error or Bad Request
// 		logrus.Errorf("Failed to check existing user: %s", err.Error())
// 		c.JSON(500, gin.H{
// 			"message": "Internal Server Error",
// 		})
// 		return
// 	}

// 	// Check if Email Already Exist Return Message error Bad Request
// 	if existingUser.Email != "" {
// 		common.SendErrorResponse(c, 400, "Email "+email+" already exist")
// 		return
// 	}

// 	// If Email Not Exist Create New User
// 	newUser, err := a.userUC.CreateUser(user)
// 	if err != nil {
// 		// Log For Create User Error
// 		logrus.Errorf("Failed to create user: %s", err.Error())
// 		c.JSON(500, gin.H{
// 			"message": "Internal Server Error",
// 		})
// 		return
// 	}

// 	// Log For Success
// 	logrus.Infof("User created successfully")
// 	c.JSON(201, gin.H{
// 		"code":    201,
// 		"message": "User created successfully",
// 		"data":    newUser,
// 	})
// }

func (a *UserController) UpdateUser(c *gin.Context) {
	// Get ID from URL parameter
	id := c.Param("id")

	// Validate ID
	if id == "" {
		c.JSON(400, gin.H{
			"message": "ID is required",
		})
		return
	}

	// Bind JSON request ke model User
	updatedUser := model.User{}
	err := c.ShouldBindJSON(&updatedUser)
	if err != nil {
		// Log For Bad Request

		c.JSON(400, gin.H{
			"message": "Failed to bind JSON: " + err.Error(),
		})
		return
	}

	// Validate if ID exist
	existingUser, err := a.userUC.FindUserById(id)
	if err != nil {
		// Log if ID not found or Error

		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// Validate If ID notfound
	if err != nil {
		common.SendErrorResponse(c, 404, "User with id "+id+" not found")
		return
	}

	// Validate If New Email same with another User
	if updatedUser.Email != existingUser.Email {
		existingUserByEmail, err := a.userUC.FindUserByEmail(updatedUser.Email)
		if err != nil {
			// Log if Email not found or Error

			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		// Validate If Email already using by another user
		if existingUser.Email != "" {
			common.SendErrorResponse(c, 400, " Email "+existingUserByEmail.Email+" already exist")
			return
		}
	}

	// Update User
	updatedUser.Id = existingUser.Id
	updatedUser, err = a.userUC.UpdateUser(updatedUser)
	if err != nil {
		// Log For Update User Error

		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// Log For Success

	c.JSON(200, gin.H{
		"code":    200,
		"message": "User updated successfully",
		"data":    updatedUser,
	})
}

func (a *UserController) DeleteUser(c *gin.Context) {
	// Get ID from URL parameter
	id := c.Param("id")

	// Validate ID
	if id == "" {
		common.SendErrorResponse(c, 400, "ID is required")
		return
	}

	// Delete User
	err := a.userUC.DeleteUser(id)
	if err != nil {
		common.SendErrorResponse(c, 500, fmt.Sprintf("Failed to delete user: %s", err))
		return
	}

	// If User Successfully Deleted

	c.JSON(200, gin.H{
		"code":    200,
		"message": "User deleted successfully",
	})
}
