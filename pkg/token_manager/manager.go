package token_manager

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"os"
	"time"
)

// eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM3ODU5MDgsImlhdCI6MTY5Mzc2NDMwOCwidXNlcklkIjoiZTQwZDRiOGMtMjg0NC00MjBlLThkZmMtN2Q5YTVmNGZhOTEwIn0.GoTjcytjdOUfMxnsoh0FW4TTjbhWgssgrmsIsij5E8NTgVEPZSAHQgV9PjBRsV14PQouKK5gLF2E9cz-s0-flPFMx4ArTCONRBPARnHulevwnvh0jE5kREcxGxFhD4Gmctalb4Nq7FBKiITNqX99esw1h-dJoHo5iA382y81KeqKOo6Sa_MxCRgjnUdh5jZNqRENqdGn5H3HiPV-to7D087tuaibyT0GET2jbnH--qdP1f_ZLzj7Fw3sNyAVbaVQp0fNmXtbGnN807ACzaNKb64jhKcni6SBrs8LAhz6KD8hi4Y1MXhOUVsrmI_1JQozNu11PYJfv0oeZVjAqFJhdc6d4k3Xpe-4QD8htABk90dle3bpBC_m0uz0xSXVZr7U78I8d4n-L-P0fjjoUuluB9M9hJR56f9Qb2rgnt9kcL6KVPvuYsvqu-5MliXbPsRmUqUFLDyAindGb3MadKGwGkzctjjf_8FQpvmuPphXSJFR5LMpNYgtKFnL40AFgqBN

type TokenConfig struct {
	AccessTokenTtl  int    `yaml:"access_token_ttl"`
	RefreshTokenTtl int    `yaml:"refresh_token_ttl"`
	Issuer          string `yaml:"issuer"`
	SigningKeyPath  string `env:"SIGNING_KEY_PATH"`
	PrivateKey      *rsa.PrivateKey
}

type tokenManager struct {
	cfg *TokenConfig
}

type TokenClaims struct {
	jwt.RegisteredClaims
}

type TokenManager interface {
	NewJWT(userId string) (string, error)
	ValidateToken(accessToken string) (*TokenClaims, error)
	NewRefreshToken() (string, error)
}

func NewTokenManager(cfg *TokenConfig) TokenManager {
	f, err := os.Open(cfg.SigningKeyPath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	key, err := io.ReadAll(f)

	if err != nil {
		panic(err)
	}

	cfg.PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(key)

	if err != nil {
		panic(err)
	}

	return &tokenManager{cfg: cfg}
}

func (t *tokenManager) NewJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodPS512, jwt.RegisteredClaims{
		Issuer:    "",
		Subject:   userId,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(t.cfg.AccessTokenTtl))),
	})

	return token.SignedString(t.cfg.PrivateKey)
}

func (t *tokenManager) ValidateToken(accessToken string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSAPSS); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}

			return t.cfg.PrivateKey.PublicKey, nil
		},
		jwt.WithIssuer(t.cfg.Issuer),
	)

	if err != nil {
		//TODO: handle error
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok {
		//TODO: handle error
		return nil, errors.New("error occurred casting token claims")
	}

	return claims, nil
}

func (t *tokenManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
