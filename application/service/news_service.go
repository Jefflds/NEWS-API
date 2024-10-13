package service

import (
	"api/news/application/domain"
	"api/news/application/port/output"
	"api/news/configuration/logger"
	"api/news/configuration/rest_err"
	"fmt"
)

type newsService struct {
	newsPort output.GetNewsPort
}

func NewNewsService(newsPort output.GetNewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsService(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("INIT GET_NEWS_SERVICE FUNCTION, subject=%s, from=%s", newsDomain.Subject, newsDomain.From))

	newsDomainRespnse, err := ns.newsPort.GetNewsPort(newsDomain)

	return newsDomainRespnse, err
}
