package auth

import (
	"net/http"
	"sqzsvc/models"
	"sqzsvc/services/token"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	if _, err := user.GetUserById(userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
