package controllers

import (
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const identityKey = "identity"

func GetIdentity(c *gin.Context) (*models.Identity, bool) {
	if ident, ok := c.Get(identityKey); ok {
		return ident.(*models.Identity), true
	} else {
		return &models.Identity{}, false
	}
}

func SetIdentity(c *gin.Context, identity *models.Identity) {
	c.Set(identityKey, identity)
}

func GetFromBodyValidated(c *gin.Context, obj interface{}) error {

	var err = c.ShouldBindJSON(obj)
	if err == nil {
		err = validator.New().Struct(obj)
	}

	return err
}
