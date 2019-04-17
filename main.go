package main

import (
	"os"

	"github.com/go-kit/kit/log"

	"github.com/terrabot-tech/service-template/application"
	"github.com/terrabot-tech/service-template/models"
	"github.com/terrabot-tech/service-template/provider"
	"github.com/terrabot-tech/service-template/repository"
	"github.com/terrabot-tech/service-template/service"
)

var (
	appConfig models.Config
)

func init() {
	models.LoadConfig(&appConfig)
}

func main() {
	logger := log.With(
		log.NewJSONLogger(os.Stderr),
		"ts", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	provider := provider.NewProvider(appConfig.SQLDataBase)
	err := provider.Open()
	if err != nil {
		_ = logger.Log("err", err)
		os.Exit(1)
	}

	rep := repository.NewRepository(provider, appConfig.SQLDataBase)
	svc := service.NewService(rep)

	app := application.NewApplication(svc, appConfig.HashSum, appConfig.ServerOpt, logger)
	app.Start()
}
