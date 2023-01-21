package token

import (
	"errors"
	"fmt"
	"os"
	"sqzsvc/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const USER_ID_CLAIM = "userId"
const EMAIL_CLAIM = "email"

func GenerateToken(user *models.User) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		USER_ID_CLAIM: user.ID,
		EMAIL_CLAIM:   user.Email,
		"exp":         time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	}
	fmt.Println("claims", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println("token", token)

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	// fmt.Println("tokenString", tokenString, err)
	return tokenString, err
}

func DecodeToken(encodedToken string) (*models.Identity, error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return &models.Identity{}, err
	}

	if !token.Valid {
		return &models.Identity{}, errors.New("invalid token")
	}

	// TODO: do we really need to verify claims???
	// if err := token.Claims.Valid(); err != nil {
	// 	fmt.Println("Invalid claims", err)
	// 	return &models.Identity{}, err
	// }

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &models.Identity{}, errors.New("invalid token claims")
	}

	ident := &models.Identity{
		UserId:    uint(claims[USER_ID_CLAIM].(float64)),
		UserEmail: claims[EMAIL_CLAIM].(string),
	}
	fmt.Println("Decoded Idenoty", *ident)
	return ident, nil
}
