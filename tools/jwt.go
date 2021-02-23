package tools

import (
	"time"

	"crypto/sha256"
	"encoding/base64"

	jwt "github.com/dgrijalva/jwt-go/v4"
)

func createToken(userID int64, duration time.Duration) *jwt.Token {

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"),
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Local().Add(time.Hour * time.Duration(duration))),
			IssuedAt:  jwt.Now(),
			Subject:   "",
			Issuer:    "goswami.ru",
		})

	// Sign and get the complete encoded token as a string

	return token
}

func generateHash(value string) string {
	return base64.URLEncoding.EncodeToString(sha256.Sum256([]byte(value)))
}
