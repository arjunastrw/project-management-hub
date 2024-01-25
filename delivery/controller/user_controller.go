package controller

import (
	"fmt"

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
	a.rg.GET("api/users", a.FindAllUser)
	a.rg.GET("api/user/:id", a.FindUserById)
	a.rg.GET("api/email/:email", a.FindUserByEmail)
	a.rg.POST("api/users", a.CreateUser)
	a.rg.PUT("api/user/:id", a.UpdateUser)
	a.rg.DELETE("api/user/:id", a.DeleteUser)
}

func (a *UserController) FindAllUser(c *gin.Context) {
	id := c.Param("id")
	user, err := a.userUC.FindUserById(id)
	if err != nil {
		common.SendErrorResponse(c, 400, "tasks id "+id+" not found")
		return
	}
	// Log For Success

	// Return For Success
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Success Get Resource",
		"data":    user,
	})
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
	// Bind JSON request to User Model
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		// Log For Bad Request

		c.JSON(400, gin.H{
			"message": "Failed to bind JSON: " + err.Error(),
		})
		return
	}

	email := c.Param("email")
	// Check If Email Already Exist
	existingUser, err := a.userUC.FindUserByEmail(email)
	if err != nil {
		// Log for Checking Existing User Error or Bad Request

		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// Check if Email Already Exist Return Message error Bad Request
	if existingUser.Email != "" {
		common.SendErrorResponse(c, 400, "Email "+email+" already exist")
		return
	}

	// If Email Not Exist Create New User
	newUser, err := a.userUC.CreateUser(user)
	if err != nil {
		// Log For Create User Error

		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// Log For Success

	c.JSON(201, gin.H{
		"code":    201,
		"message": "User created successfully",
		"data":    newUser,
	})
}

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
