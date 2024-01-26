package middleware

import (
	"log"
	"net/http"
	"strings"

	"enigma.com/projectmanagementhub/shared/service"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type AutHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

// RequireToken implements AuthMiddleware.
func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var autHeader AutHeader
		if err := c.ShouldBindHeader(&autHeader); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "RequireToken.autHeader"})
			return
		}

		tokenHeader := strings.Replace(autHeader.AuthorizationHeader, "Bearer ", "", -1)

		if tokenHeader == "" {
			log.Println("RequireToken.tokenHeader")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := a.jwtService.ParseToken(tokenHeader)
		if err != nil {
			log.Printf("RequireToken.ParseToken: %v", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", claims["user_id"])

		validRole := false

		for _, role := range roles {
			if role == claims["role"] {
				validRole = true
				break
			}

		}

		if !validRole {
			log.Printf("RequireToken.validRole: %v", err.Error())
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
	}
}
