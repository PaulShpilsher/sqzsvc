package controllers

import (
	"log"
	"net/http"
	"sqzsvc/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type UrlController struct{}

// GET: /:shortCode
func (me *UrlController) GotoLongUrl(c *gin.Context) {

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
