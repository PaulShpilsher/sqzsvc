package token

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Identity struct {
	UserId    uint
	UserEmail string
}

func GenerateToken(identity *Identity) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id":    identity.UserId,
		"user_email": identity.UserEmail,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	}
	fmt.Println("claims", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println("token", token)

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	// fmt.Println("tokenString", tokenString, err)
	return tokenString, err
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	bearerTokenParts := strings.Split(bearerToken, " ")
	if len(bearerTokenParts) == 2 {
		return bearerTokenParts[1]
	}
	return ""
}

func GetIdenitiyFromToken(c *gin.Context) (*Identity, error) {

	tokenString := ExtractToken(c)
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return &Identity{}, err
	}

	if err := t.Claims.Valid(); err != nil {
		fmt.Println("Invalid claims", err)
		return &Identity{}, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return &Identity{}, errors.New("invalid token claims")
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
	if err != nil {
		return &Identity{}, err
	}

	ident := Identity{
		UserId:    uint(userId),
		UserEmail: claims["user_email"].(string),
	}
	fmt.Println(ident)
	return &ident, nil

	// return &Identity{
	// 	UserId:    int(userId),
	// 	UserEmail: claims["user_email"].(string),
	// }, nil
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("nexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	fmt.Println(t, err)

	if err != nil {
		return err
	}
	return nil
}

// func ExtractTokenID(c *gin.Context) (uint, error) {

// 	tokenString := ExtractToken(c)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("TOKEN_SECRET")), nil
// 	})
// 	if err != nil {
// 		return 0, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
// 		if err != nil {
// 			return 0, err
// 		}
// 		return uint(uid), nil
// 	}
// 	return 0, nil
// }
