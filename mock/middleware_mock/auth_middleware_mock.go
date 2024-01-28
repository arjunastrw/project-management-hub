package middleware_mock

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type UserMiddlewareMock struct {
	mock.Mock
}

func (a *UserMiddlewareMock) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			return
		}
		ctx.Next()
	}
}
