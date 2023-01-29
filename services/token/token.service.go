package token

import (
	"errors"
	"fmt"
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

func DecodeToken(encodedToken string) (*Identity, error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.TokenSecret), nil
	})
	if err != nil {
		return &Identity{}, err
	}

	if !token.Valid {
		return &Identity{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &Identity{}, errors.New("invalid token claims")
	}

	ident := &Identity{
		UserID:    uint(claims[userIdClaimKey].(float64)),
		UserEmail: claims[emailClaimKey].(string),
	}
	return ident, nil
}
