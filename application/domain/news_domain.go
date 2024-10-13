package domain


type NewsReqDomain struct {
	Subject string
	From    string
}

type NewsDomain struct {
	Status       string
	TotalResults int
	Articles     []Article
}

type Article struct {
	Source Source
	Author string
	Description string
	URL string
	UrlToImage string
	PublishedAt string
	Content string
}

type Source struct {
	Id   string
	Name string
}
