package news_http

import (
	"api/news/adapter/output/model/response"
	"api/news/application/domain"
	"api/news/configuration/env"
	"api/news/configuration/rest_err"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type NewsClient struct{}

func NewNewsClient() *NewsClient {
	client = resty.New().SetBaseURL(env.GetNewsURL())
	return &NewsClient{}
}

var (
	client *resty.Client
)

func (nc *NewsClient) GetNewsPort(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {

	newsReponse := &response.NewsClientResponse{}

	_, err := client.R().SetQueryParams(map[string]string{
		"q":      newsDomain.Subject,
		"from":   newsDomain.From,
		"apiKey": env.GetNewsTokenAPI(),
	}).SetResult(newsReponse).Get("/everything")

	fmt.Println(err)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error to trying to call NewsAPI with Params")
	}

	newsResponseDomain := &domain.NewsDomain{}
	copier.Copy(newsResponseDomain, newsReponse)

	return newsResponseDomain, nil
}
