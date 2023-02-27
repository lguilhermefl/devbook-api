package authentication

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"userId": userID,
	}
	// permissions["authorized"] = true
	// permissions["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.JWTScecret)
}