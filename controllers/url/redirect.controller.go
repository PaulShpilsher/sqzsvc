package url

import (
	"log"
	"net/http"
	urlService "sqzsvc/services/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// GET: /:shortCode
func RedirectShortCode(c *gin.Context) {

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
