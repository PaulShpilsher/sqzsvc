package url

import (
	"net/http"
	"net/url"
	"sqzsvc/controllers"
	urlService "sqzsvc/services/url"

	"github.com/gin-gonic/gin"
)

// /////////
type LongUrlInput struct {
	Url string `json:"url" binding:"required" validate:"required,url"`
}

// POST: /api/short-code
func CreateShortCode(c *gin.Context) {
	identity, ok := controllers.GetIdentity(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unable to get indentity"})
		return
	}

	input := &LongUrlInput{}
	if err := controllers.GetFromBodyValidated(c, input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := url.ParseRequestURI(input.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if shortCode, err := urlService.RegisterLongUrl(identity, url.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"shortCode": shortCode})
	}
}
