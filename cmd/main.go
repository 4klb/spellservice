package main

import (
	"main/internal/services/api"
	"main/internal/services/routes"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		println("не  удалось создать логер")
		return
	}
	log := logger.Sugar()
	defer log.Sync()

	text, err := api.Tmp(log)
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
