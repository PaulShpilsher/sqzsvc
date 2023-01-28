package controllers

import (
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
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
