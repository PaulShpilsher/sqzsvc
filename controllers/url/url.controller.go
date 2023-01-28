package url

import (
	"log"
	"net/http"
	"net/url"
	"sqzsvc/controllers"
	"sqzsvc/services"
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

	service := &services.ShortCodeService{}
	if longUrl, err := service.GetLongUrl(shortCode); err == nil {
		c.Redirect(http.StatusFound, longUrl)
	} else {
		log.Println("Failed to get long url: ", err)
		c.Status(http.StatusNotFound)
	}

	// host := c.GetHeader("Host")
	// ip := c.ClientIP()
	// log.Printf("Short code '%s', host '%s', ip '%s'", shortCode, host, ip)

	// location, _ := url.Parse("http://www.cnn.com")
	// log.Println(location)
	// c.Redirect(http.StatusFound, "http://www.cnn.com")
}

///////////

type RegisterLongUrlInput struct {
	Url string `json:"url" binding:"required"`
}

// POST: /api/short-code
func CreateShortCode(c *gin.Context) {
	ident, ok := controllers.GetIdentity(c)
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

	service := &services.ShortCodeService{Identity: ident}
	if shortCode, err := service.RegisterLongUrl(url.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"shortCode": shortCode})
	}
}
