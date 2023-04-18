package token_manager

import "time"

type Option func(t *tokenManager)

func WithSigningKey(signingKey string) Option {
	return func(t *tokenManager) {
		t.signingKey = signingKey
	}
}

func WithTTL(ttl time.Duration) Option {
	return func(t *tokenManager) {
		t.ttl = ttl
	}
}
