package token

import (
	"errors"
	"fmt"
	"log"
	"sqzsvc/models"
	"sqzsvc/services/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	userIdClaimKey = "userId"
	emailClaimKey  = "email"
)

func GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		userIdClaimKey: user.ID,
		emailClaimKey:  user.Email,
		"exp":          time.Now().Add(time.Hour * time.Duration(config.TokenHourLifespan)).Unix(),
	}
	// log.Println("claims", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println("token", token)

	tokenString, err := token.SignedString([]byte(config.TokenSecret))
	// fmt.Println("tokenString", tokenString, err)
	return tokenString, err
}

func DecodeToken(encodedToken string) (*models.Identity, error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.TokenSecret), nil
	})
	if err != nil {
		return &models.Identity{}, err
	}

	if !token.Valid {
		return &models.Identity{}, errors.New("invalid token")
	}

	// QUESTION: do we really need to verify claims???
	// if err := token.Claims.Valid(); err != nil {
	// 	fmt.Println("Invalid claims", err)
	// 	return &models.Identity{}, err
	// }

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &models.Identity{}, errors.New("invalid token claims")
	}

	ident := &models.Identity{
		UserID:    uint(claims[userIdClaimKey].(float64)),
		UserEmail: claims[emailClaimKey].(string),
	}
	log.Println("Decoded Idenoty", *ident)
	return ident, nil
}
