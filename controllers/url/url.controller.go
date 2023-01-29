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

	if shortCode, err := urlService.SubmitUrl(url.String(), c.ClientIP()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"shortCode": shortCode})
	}
}
