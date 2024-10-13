package routes

import (
	"api/news/adapter/input/controller"
	"api/news/adapter/output/news_http"
	"api/news/application/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsClient := news_http.NewNewsClient()
	newsServices := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsServices)

	r.GET("/news", newsController.GetNews)
}
