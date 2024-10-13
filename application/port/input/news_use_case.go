package input

import (
	"api/news/application/domain"
	"api/news/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
