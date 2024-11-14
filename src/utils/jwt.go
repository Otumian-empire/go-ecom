package utils

import (
	"encoding/base64"
	"fmt"
	"otumian-empire/go-ecom/src/config"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authCustomClaims struct {
	Id config.IdType `json:"id"`
	jwt.RegisteredClaims
}

// GenerateJwt generates a JWT with custom claims and a specified TTL.
func GenerateJwt(id config.IdType) (string, error) {
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
		return "", err
	}

	return authToken, nil
}

func VerifyJwt(encodedToken string) (*jwt.Token, error) {
	environ, _ := config.GetEnvirons()

	payload, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}

		return []byte(environ.JwtSecret), nil
	})

	return payload, err
}

func IsExpiredToken(claims jwt.MapClaims) bool {
	// Get the expiration time from the "exp" claim
	expirationTime, ok := claims["exp"].(float64)
	if !ok {
		return true
	}

	// Convert the expiration time to a Unix timestamp in seconds
	expirationUnix := int64(expirationTime)

	// Get the current time
	currentTime := time.Now().Unix()

	return currentTime >= expirationUnix
}

func IsValidIssuer(claims jwt.MapClaims, issuer string) bool {
	// check the issuer
	_issuer, ok := claims["iss"].(string)
	if !ok {
		return true
	}

	return issuer == _issuer
}

func GetIdFromClaim(claims jwt.MapClaims) (config.IdType, error) {
	idClaim, ok := claims["id"].(float64)

	if !ok {
		return -1, fmt.Errorf("some sort of error occurred")
	}

	return config.IdType(idClaim), nil
}

func ValidateApiKey(token string) bool {
	// Split the token by space
	authorization := strings.Split(token, " ")
	if len(authorization) != 2 {
		return false
	}

	// Check if the keyword is "Bearer"
	bearerKeyword, encodedToken := authorization[0], authorization[1]
	if bearerKeyword != "Bearer" {
		return false
	}

	// Split the token into header, payload, and signature
	parts := strings.Split(encodedToken, ".")
	if len(parts) != 3 {
		return false
	}

	// Decode header
	_, headerErr := base64.RawURLEncoding.DecodeString(parts[0])
	if headerErr != nil {
		return false
	}

	// Decode payload
	_, payloadErr := base64.RawURLEncoding.DecodeString(parts[1])
	if payloadErr != nil {
		return false
	}

	// Basic check: The signature should be a valid base64 URL string.
	_, signatureErr := base64.RawURLEncoding.DecodeString(parts[2])

	return signatureErr == nil
}
