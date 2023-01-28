package url

import (
	"log"
	"net/http"
	"net/url"
	"sqzsvc/controllers"
	urlService "sqzsvc/services/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// GET: /:shortCode
func GotoLongUrl(c *gin.Context) {

	shortCode := strings.TrimSpace(c.Param("shortCode"))
	if len(shortCode) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	if longUrl, err := urlService.GetLongUrl(shortCode); err == nil {
		// TODO: Log click with client IP ip := c.ClientIP()
		c.Redirect(http.StatusFound, longUrl)
	} else {
		log.Println("Failed to get long url: ", err)
		c.Status(http.StatusNotFound)
	}
}

///////////

type RegisterLongUrlInput struct {
	Url string `json:"url" binding:"required"`
}

// POST: /api/short-code
func CreateShortCode(c *gin.Context) {
	identity, ok := controllers.GetIdentity(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unable to get indentity"})
		return
	}

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

	if shortCode, err := urlService.RegisterLongUrl(identity, url.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"shortCode": shortCode})
	}
}
