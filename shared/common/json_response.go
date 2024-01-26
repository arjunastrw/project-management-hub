package common

import (
	"net/http"

	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/gin-gonic/gin"
)

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, shared_model.Status{
		Code:    code,
		Message: message,
	})
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, shared_model.SingleResponse{
		Status: shared_model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendPagedResponse(ctx *gin.Context, data []interface{}, page shared_model.Paging, message string) {
	ctx.JSON(http.StatusOK, shared_model.PagedResponse{
		Status: shared_model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data:   data,
		Paging: page,
	})
}

func SendSuccesResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, shared_model.StatusSucces{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendCreatedResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, shared_model.SingleResponse{
		Status: shared_model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}
