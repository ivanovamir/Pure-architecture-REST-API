package app

import (
	"context"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/config"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/transport/http/handler"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/logger"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/password_manager"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/postgresql"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/server"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
)

func Run() {
	cfg := config.NewConfig()

	if cfg == nil {
		// TODO error handler
		log.Println("error occurred loading config")
		return
	}

	lg := logger.NewLogger(logger.WithCfg(&cfg.LoggerConfig), logger.WithAppVersion(cfg.AppVersion.AppVersion))

	if lg == nil {
		// TODO error handler
		log.Println("error occurred loading logger")
		return
	}

	//db config
	db, err := postgresql.NewPostgresDB(context.Background(), &cfg.PostgresDBConfig)

	if err != nil {
		// TODO error handler
		log.Println("error occurred connection to postgresql")
		return
	}

	defer db.Close()

	passwordManager := password_manager.NewPasswordManager(&cfg.PasswordManagerConfig)

	//tokenManager := token_manager.NewTokenManager(&cfg.TokenConfig)

	httpMx := httprouter.New()

	repo := repository.NewRepository(db)

	service := service.NewService(lg, repo, passwordManager)

	httpHandler := handler.NewHttpHandler(httpMx, service)

	httpHandler.Router()

	ln, err := net.Listen(cfg.HandlerConfig.ListenType, fmt.Sprintf("%s:%s", cfg.HandlerConfig.ListenAddr, cfg.HandlerConfig.ListenPort))

	if err != nil {
		lg.Error(err.Error())
		return
	}

	srv := server.NewServer(
		server.WithSrv(&http.Server{
			Addr:    fmt.Sprintf(":%s", cfg.HandlerConfig.ListenPort),
			Handler: httpMx,
		}),
		server.WithListener(&ln),
	)

	quiteCh := make(chan struct{})
	go func() {
		if err = srv.Run(); err != nil {
			//TODO: handle error
			lg.Error(err.Error())
			quiteCh <- struct{}{}
			return
		}
	}()

	<-quiteCh
}
