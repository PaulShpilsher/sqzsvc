package middlewares

import (
	"net/http"
	"sqzsvc/services/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ident, err := token.GetIdenitiyFromToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Set("identity", ident)

		// err := token.TokenValid(c)
		// if err != nil {
		// 	c.String(http.StatusUnauthorized, "Unauthorized")
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}
