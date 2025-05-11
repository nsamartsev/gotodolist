package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func createToken(username string) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject (user identifier)
		"iss": "todo-app",                       // Issuer
		"aud": getRole(username),                // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})
	// Print information about the created token
	fmt.Printf("Token claims added: %+v\n", claims)

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getRole(username string) string {
	if username == "senior" {
		return "senior"
	}
	return "employee"
}

// Function to verify JWT tokens
func verifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}
