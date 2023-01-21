package controllers

import (
	"sqzsvc/models"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (controller Controller) GetIdenity(c *gin.Context) (*models.Identity, bool) {
	if ident, ok := c.Get("identity"); ok {
		return ident.(*models.Identity), true
	} else {
		return &models.Identity{}, false
	}
}
