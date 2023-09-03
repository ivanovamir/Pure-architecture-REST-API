package password_manager

import (
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

type PasswordConfig struct {
	TimeHash    uint32 `yaml:"time_hash"`
	Memory      uint32 `yaml:"memory"`
	ThreadCount uint8  `yaml:"thread_count"`
	KeyLen      uint32 `yaml:"key_len"`
	SaltLength  int    `yaml:"salt_length"`
}

type passwordManager struct {
	cfg *PasswordConfig
}

// TODO: add password validator
type PasswordManager interface {
	PasswordHash(password, salt []byte) []byte
	Salt() []byte
}

func NewPasswordManager(cfg *PasswordConfig) PasswordManager {
	return &passwordManager{cfg: cfg}
}

func (p *passwordManager) Salt() []byte {
	var salt = make([]byte, p.cfg.SaltLength)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}

	return salt
}

func (p *passwordManager) PasswordHash(password, salt []byte) []byte {
	return argon2.IDKey(password, salt, p.cfg.TimeHash, p.cfg.Memory, p.cfg.ThreadCount, p.cfg.KeyLen)
}
