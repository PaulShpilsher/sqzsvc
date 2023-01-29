package url

import (
	"log"
	"net/http"
	transitionService "sqzsvc/services/transition"
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
		transitionService.LogTransition(shortCode, c.ClientIP())
		c.Redirect(http.StatusFound, url) // in production use c.Redirect(http.StatusMovedPermanently, url)
	} else {
		log.Printf("Sort Code [%s] failed produce url | %s\n", shortCode, err.Error())
		c.Status(http.StatusNotFound)
	}
}
