package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type Payload struct {
	ClientId  string    `json:"client_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func CreateToken(ClientId, SecretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client_id": ClientId,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Token generated successfully")
		return "", err
	}
	return tokenString, nil
}

// 1. Parse and validate the JWT token
// 2. Check if token is expired
// 3. Return the claims if valid
func VerifyToken(tokenString string, secretKey string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error().
				Str("signing_method", token.Header["alg"].(string)).
				Msg("Invalid signing method")
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to parse token")
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token is expired
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				log.Error().
					Msg("Token is expired")
				return nil, jwt.ErrTokenExpired
			}
		}
		return &claims, nil
	}

	log.Error().
		Msg("Invalid token claims")
	return nil, jwt.ErrInvalidKey
}
