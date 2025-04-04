package utils

import (
	"fmt"
	"log"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func getJwtSecret() []byte {
	jwtSecret := viper.GetString("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in environment variables")
	}
	return []byte(jwtSecret)
}
func GetToken(email string) (string, error) {

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp"  :  time.Now().Add(time.Hour * 1).Unix(),
	})
	
	jwtSecret := getJwtSecret()
	
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (string, error) {
	jwtSecret := getJwtSecret()
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}
	username := claims["email"].(string)
	return username, nil
}