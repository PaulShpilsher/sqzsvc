package auth

import (
	"net/http"
	"sqzsvc/controllers"
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	ident, ok := controllers.Controller{}.GetIdenity(c)
	if !ok {
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
