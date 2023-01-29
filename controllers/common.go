package controllers

import (
	"sqzsvc/services/token"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const identityKey = "identity"

func GetIdentity(c *gin.Context) (*token.Identity, bool) {
	if ident, ok := c.Get(identityKey); ok {
		return ident.(*token.Identity), true
	} else {
		return &token.Identity{}, false
	}
}

func SetIdentity(c *gin.Context, identity *token.Identity) {
	c.Set(identityKey, identity)
}

func GetFromBodyValidated(c *gin.Context, obj interface{}) error {

	var err = c.ShouldBindJSON(obj)
	if err == nil {
		err = validator.New().Struct(obj)
	}

	return err
}
