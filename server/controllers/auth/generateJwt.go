package auth

import (
	"time"
	"os"

	"github.com/NoamBoni/bitly/server/helpers"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	Id               int `json:"id"`
}

func GenerateJWT(id int) (string, error) {
	helpers.LoadEnv()
	JWTsecret := []byte(os.Getenv("JWT_SECRET"))
	TokenDuration := time.Hour * 3
	claims := Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenDuration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer:    "bitly",
		},
		id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTsecret)
}
