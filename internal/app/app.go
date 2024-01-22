package app

import (
	"errors"
	"net/http"
	"userLoggingProject/internal/core/config"
	"userLoggingProject/internal/core/service_provider"
	"userLoggingProject/internal/features/logs/repository"
	"userLoggingProject/internal/features/logs/service"
	"userLoggingProject/pkg/logger"
	"userLoggingProject/transport/rest"
)

type App struct {
}

// @title Logger Service api
// @version 1.0
// @description REST API for Logger Service

// @host localhost:8080
// @BasePath /api/v1/

func RunApp() {

	conf, err := config.Read()
	if err != nil {
		logger.Error(err)
		return
	}

	logRepo := repository.NewMemoryRepo()

	logService := service.NewService(logRepo)

	services := service_provider.NewServiceProvider(logService)

	handlers := rest.NewHandler(services)

	switch conf.LaunchType {
	case config.REST:
		serv := rest.NewServer(&conf, handlers.Init(&conf))
		func() {
			if err := serv.Run(); !errors.Is(err, http.ErrServerClosed) {
				logger.Errorf("error running server: %s\n", err.Error())
			}
		}()
		logger.Info("Server has been started")

		if err != nil {
			logger.Error(err)
			return
		}
	default:
		logger.Info("Launch Type not set, shutDown")
		return
	}

}
