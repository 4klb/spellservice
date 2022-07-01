package main

import (
	"main/config"
	"main/internal/services/api"
	"main/internal/services/routes"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		println("could not create a logger")
		return
	}
	log := logger.Sugar()
	defer log.Sync()

	v := config.ParseConfig()
	if v == nil {
		log.Debug("")
		return
	}

	text, err := api.GetText(log, v)
	if err != nil {
		log.Debug(err.Error())
		return
	}

	responces, err := api.GetResponce(text.Texts, log, v)
	if err != nil {
		log.Debug(err)
		return
	}

	api.Replace(responces, text.Texts)
	err = routes.SetUpRoutes(text, log, v)
	if err != nil {
		log.Debug(err)
		return
	}
}
