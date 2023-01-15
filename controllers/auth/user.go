package auth

import (
	"net/http"
	"sqzsvc/models"
	"sqzsvc/services/token"

	"github.com/gin-gonic/gin"
)

func GetIdenity(c *gin.Context) *token.Identity {
	ident, ok := c.Get("identity")
	if ok {
		return ident.(*token.Identity)
	} else {
		return nil
	}
}

func CurrentUser(c *gin.Context) {
	ident := GetIdenity(c)
	if ident == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get indentity"})
		return
	}

	user := &models.User{}
	if _, err := user.GetUserById(ident.UserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
