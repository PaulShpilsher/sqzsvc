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
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing short code"})
		return
	}

	if url, err := urlService.GetUrl(shortCode); err == nil {
		// TODO: Log click with client IP ip := c.ClientIP()
		c.Redirect(http.StatusFound, url)
	} else {
		log.Printf("Sort Code [%s] failed produce url | %s\n", shortCode, err.Error())
		c.Status(http.StatusNotFound)
	}
}
