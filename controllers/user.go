package controllers

import (
	"log"
	"net/http"
	"sqzsvc/models"
	"sqzsvc/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*Controller
}

func (me *UserController) CurrentUser(c *gin.Context) {
	ident, ok := me.GetIdenity(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to indentity user"})
		return
	}

	user := &models.User{}
	if _, err := user.GetUserById(ident.UserId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

///////////  Register new user

type RegisterLongUrlInput struct {
	Url string `json:"url" binding:"required"`
}

func (me *UserController) RegisterLongUrl(c *gin.Context) {
	ident, ok := me.GetIdenity(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unable to get indentity"})
		return
	}

	log.Println(ident)

	var input RegisterLongUrlInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := &services.ShortCodeService{
		Identity: ident,
	}

	if shortCode, err := service.RegisterLongUrl(input.Url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		// TODO: construct short url

		c.JSON(http.StatusOK, gin.H{"shortCode": shortCode, "shortUrl": "http://localhost:5555/" + shortCode})
	}

}
