package main

import (
	"api/news/adapter/input/routes"
	"api/news/configuration/env"
	"api/news/configuration/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("GETTING DOTENV")
	env.Init()

	logger.Info("ABOUT TO START APPLICATION")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error trying to init application on port 8080", err)
		return
	}
}
