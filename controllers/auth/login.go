package auth

import (
	"net/http"
	"sqzsvc/models"
	"sqzsvc/services"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	if _, err := user.GetUserByEmail(input.Email); err != nil {
		// user not found
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
		return
	}

	if err := services.VerifyPassword(input.Password, user.Password); err != nil {
		// password mismatch
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
		return
	}

	token, err := services.GenerateToken(user.ID)
	if err != nil {
		// cant generate token
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
