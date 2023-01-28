package auth

import (
	"log"
	"net/http"
	"sqzsvc/models"
	"sqzsvc/services/token"
	"sqzsvc/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

///////////  Register new user

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`    // max 256 chars
	Password string `json:"password" binding:"required"` // max 72 chars
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: validate input

	user := models.User{
		Email:    strings.TrimSpace(input.Email),
		Password: strings.TrimSpace(input.Password),
	}

	_, err := user.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("User registered", user)

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

///////////  Login user

type LoginInput struct {
	Email    string `json:"email" binding:"required"`    // max 256 chars
	Password string `json:"password" binding:"required"` // max 72 chars
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	if _, ok := user.GetUserByEmail(input.Email); !ok {
		// user not found
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
		return
	}

	if err := utils.VerifyPassword(input.Password, user.Password); err != nil {
		// password mismatch
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
		return
	}

	if token, err := token.GenerateToken(user); err != nil {
		// cant generate token
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
