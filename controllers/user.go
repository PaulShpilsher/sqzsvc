package controllers

import (
	"log"
	"net/http"
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*Controller
}

func (me *UserController) CurrentUser(c *gin.Context) {
	ident, ok := me.GetIdenity(c)
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

func (me *UserController) RegosterLongUrl(c *gin.Context) {
	ident, ok := me.GetIdenity(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get indentity"})
		return
	}

	log.Println(ident)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": ident})

}
