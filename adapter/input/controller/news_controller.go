package controller

import (
	"api/news/adapter/input/model/request"
	"api/news/application/domain"
	"api/news/application/port/input"
	"api/news/configuration/logger"
	"api/news/configuration/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(
	newsUseCase input.NewsUseCase,
) *newsController {
	return &newsController{
		newsUseCase: newsUseCase,
	}
}

func (nc *newsController) GetNews(c *gin.Context) {

	logger.Info("INIT GETNEWS CONTROLLER API")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to  validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	newsDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From:    newsRequest.From.Format("2006-01-02"),
	}

	newsReponseDomain, err := nc.newsUseCase.GetNewsService(newsDomain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, newsReponseDomain)
}
