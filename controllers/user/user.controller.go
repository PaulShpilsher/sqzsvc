package user

import (
	"net/http"
	"sqzsvc/controllers"
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	ident, ok := controllers.GetIdentity(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to indentity user"})
		return
	}

	user := &models.User{}
	if _, err := user.GetUserById(ident.UserID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
