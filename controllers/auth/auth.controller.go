package auth

import (
	"log"
	"net/http"
	"sqzsvc/controllers"
	"sqzsvc/models"
	"sqzsvc/services/token"
	"sqzsvc/utils"

	"github.com/gin-gonic/gin"
)

///////////  Register new user

func Register(c *gin.Context) {

	input := &CredentialsInput{}
	if err := controllers.GetFromBodyValidated(c, input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: apply password policy validation (e.g. does not contain whitespaces, etc..)

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	if _, err := user.SaveUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("User registered", user)
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

///////////  Login user

func Login(c *gin.Context) {

	input := &CredentialsInput{}
	if err := controllers.GetFromBodyValidated(c, input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	if _, ok := user.GetUserByEmail(input.Email); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
		return
	}

	if err := utils.VerifyPassword(input.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
		return
	}

	if token, err := token.GenerateToken(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, &AuthOutput{Token: token})
	}
}
