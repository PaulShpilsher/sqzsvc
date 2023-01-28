package token

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sqzsvc/models"
	"strconv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	userIdClaimKey = "userId"
	emailClaimKey  = "email"
)

var (
	signatureKey []byte = nil
	once         sync.Once
)

func getSignatureKey() []byte {
	once.Do(func() {
		signatureKey = []byte(os.Getenv("TOKEN_SECRET"))
	})
	return signatureKey
}

func GenerateToken(user *models.User) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		userIdClaimKey: user.ID,
		emailClaimKey:  user.Email,
		"exp":          time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	}
	log.Println("claims", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println("token", token)

	tokenString, err := token.SignedString(getSignatureKey())
	// fmt.Println("tokenString", tokenString, err)
	return tokenString, err
}

func DecodeToken(encodedToken string) (*models.Identity, error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getSignatureKey(), nil
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
		UserID:    uint(claims[userIdClaimKey].(float64)),
		UserEmail: claims[emailClaimKey].(string),
	}
	log.Println("Decoded Idenoty", *ident)
	return ident, nil
}
