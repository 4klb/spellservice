package main

import (
	"log"
	"main/internal/services/api"
	"main/internal/services/routes"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Println("could not create a logger")
		return
	}
	log := logger.Sugar()
	defer log.Sync()

	text, err := api.GetText(log)
	if err != nil {
		log.Debug(err.Error())
		return
	}

	err = routes.SetUpRoutes(text, log)
	if err != nil {
		log.Debug(err)
		return
	}
}
