package middlewares

import (
	"log"
	"net/http"
	"sqzsvc/controllers"
	tokenService "sqzsvc/services/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if encodedToken := extractToken(c); encodedToken != "" {
			if identity, err := tokenService.DecodeToken(encodedToken); err == nil {
				controllers.SetIdentity(c, identity)
				c.Next()
				return
			} else {
				log.Println("Token decoding failed", err)
			}
		}

		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
	}
}

func extractToken(c *gin.Context) string {
	if authorizationHeader := c.Request.Header.Get("Authorization"); authorizationHeader != "" {
		if bearerTokenParts := strings.Split(authorizationHeader, " "); len(bearerTokenParts) == 2 {
			return bearerTokenParts[1]
		}
	}
	return ""
}
