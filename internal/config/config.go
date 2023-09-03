package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/logger"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/password_manager"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/postgresql"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/token_manager"
	"github.com/joho/godotenv"
	"os"
)

type configOptions func(*Config)

type Config struct {
	AppVersion            AppConfig                       `yaml:"app"`
	HandlerConfig         ServerConfig                    `yaml:"server"`
	LoggerConfig          logger.LoggerConfig             `yaml:"logger"`
	PostgresDBConfig      postgresql.PostgresDBConfig     `yaml:"postgres"`
	TokenConfig           token_manager.TokenConfig       `yaml:"token"`
	PasswordManagerConfig password_manager.PasswordConfig `yaml:"password"`
}

var globalConfigOptions []configOptions

func addGlobalConfigOption(opt configOptions) {
	globalConfigOptions = append(globalConfigOptions, opt)
}

func NewConfig() *Config {
	cfg := new(Config)

	if err := getEnv(); err != nil {
		return nil
	}

	validateServerMode()

	if err := cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), cfg); err != nil {
		return nil
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil
	}

	for _, opt := range globalConfigOptions {
		opt(cfg)
	}

	return cfg
}

func getEnv() error {
	return godotenv.Load(".env")
}
