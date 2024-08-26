package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github/losunioncode/library-managment-system/internal/database"
	"time"
)

var jwtKey = []byte(database.GetEnvDB()[2])

type JWTClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (tokenString string, err error) {
	expiritaionTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiritaionTime),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = tokenClaims.SignedString(jwtKey)

	return

}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return
	}
	return
}
