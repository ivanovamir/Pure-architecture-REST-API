package token_manager

/* This package is designed to work with JWT tokens. The structure of the package (tokenManager) includes:

# type tokenManager struct {
# 	signingKey string # unique secret signing key
# 	ttl        time.Duration # Token time to live
# }

This package uses the "crypto/rand" package, it's better, than "math/rand", which is a pseudo-random generation.
If there is a need to use extremely precise and maximally unique values,
"crypto/rand" package should be used.

Refresh token should preferably be stored in a Redis cache,
with user id (uuid) as the key and refresh token itself as the value.
*/

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type tokenManager struct {
	signingKey string
	ttl        time.Duration
}

type TokenManager interface {
	NewJWT(userId string) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

func NewTokenManager(signingKey string, ttl time.Duration) TokenManager {
	return &tokenManager{
		signingKey: signingKey,
		ttl:        ttl,
	}
}

func (t *tokenManager) NewJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.ttl) * time.Hour)),
	})

	return token.SignedString([]byte(t.signingKey))
}

func (t *tokenManager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(t.signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", fmt.Errorf("error occured getting user claims from token")
	}

	/* Get "sub" object from map of jwt payload */
	return claims["sub"].(string), nil
}

func (t *tokenManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
