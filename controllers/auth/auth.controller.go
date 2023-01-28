package auth

import (
	"log"
	"net/http"
	"sqzsvc/models"
	"sqzsvc/services/token"
	"sqzsvc/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

///////////  Register new user

func Register(c *gin.Context) {

	credentials, err := getCredentials(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    strings.TrimSpace(credentials.Email),
		Password: strings.TrimSpace(credentials.Password),
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

	credentials, err := getCredentials(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	if _, ok := user.GetUserByEmail(credentials.Email); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
		return
	}

	err = utils.VerifyPassword(credentials.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
		return
	}

	if token, err := token.GenerateToken(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, &AuthOutput{Token: token})
	}
}

func getCredentials(c *gin.Context) (*CredentialsInput, error) {

	var credentials = &CredentialsInput{}
	var err = c.ShouldBindJSON(credentials)
	if err == nil {
		err = validator.New().Struct(credentials)
	}

	return credentials, err
}
