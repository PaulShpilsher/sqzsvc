package controllers

import (
	"log"
	"net/http"
	"net/url"
	"sqzsvc/models"
	"sqzsvc/services"
	"strings"

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
	if _, err := user.GetUserById(ident.UserID); err != nil {
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

	url, err := url.ParseRequestURI(strings.TrimSpace(input.Url))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := &services.ShortCodeService{Identity: ident}
	if shortCode, err := service.RegisterLongUrl(url.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"shortCode": shortCode})
	}
}
