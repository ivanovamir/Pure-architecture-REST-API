package app

import (
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/server"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/transport/handler"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/cache"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/postgresql"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/token_manager"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func initConfig() error {
	viper.AddConfigPath("cmd/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func Run() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured load config file: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error occured load .env file: %s", err.Error())
	}

	router := httprouter.New()

	//db config
	db, err := postgresql.NewPostgresDB(&postgresql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.db_name"),
		SslMode:  viper.GetString("db.ssl_mode"),
	})

	if err != nil {
		log.Fatalf("error occured db: %s", err.Error())
	}

	accessTokenTtl, err := time.ParseDuration(viper.GetString("token.access_token_ttl"))
	refreshTokenTtl, err := time.ParseDuration(viper.GetString("token.refresh_token_ttl"))

	if err != nil {
		return
	}

	tokenManager := token_manager.NewTokenManager(
		token_manager.WithSigningKey(os.Getenv("SIGNED_KEY")),
		token_manager.WithTTL(accessTokenTtl),
	)

	cacheClient := cache.NewRedisClient(
		cache.WithAddress(viper.GetString("redis.address")),
		cache.WithPassword(os.Getenv("REDIS_DB_PASSWORD")),
		cache.WithDB(viper.GetInt("redis.token_db")),
	)

	// Entities
	repository := repository.NewRepository(db, cacheClient)

	// Use cases
	service := service.NewService(
		repository,
		tokenManager,
		refreshTokenTtl,
	)

	// Gateway
	handler := handler.NewHttpHandler(router, service)

	// Register router
	handler.Router()

	listener, err := net.Listen(viper.GetString("server.port"), fmt.Sprintf(":%s", viper.GetString("server.port")))

	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	// Config server
	srv := server.NewServer(&http.Server{
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}, listener)

	// Run server
	if err = srv.Run(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
