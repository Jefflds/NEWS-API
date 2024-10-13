package output

import (
	"api/news/application/domain"
	"api/news/configuration/rest_err"
)

type GetNewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
