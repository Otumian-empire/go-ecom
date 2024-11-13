package utils

import (
	"fmt"
	"otumian-empire/go-ecom/src/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authCustomClaims struct {
	Id config.IdType `json:"id"`
	jwt.RegisteredClaims
}

// GenerateJwt generates a JWT with custom claims and a specified TTL.
func GenerateJwt(id config.IdType) string {
	environ, _ := config.GetEnvirons()

	// Parse TTL from environment
	ttl, _ := strconv.Atoi(environ.JwtTtl)

	// Calculate the expiration time
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(ttl)))

	// Define custom claims
	claims := &authCustomClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    environ.JwtIssuer,
		},
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	authToken, err := token.SignedString([]byte(environ.JwtSecret))
	if err != nil {
		panic(fmt.Sprintf("Failed to sign token: %v", err))
	}

	return authToken
}

func VerifyJwt() {}
