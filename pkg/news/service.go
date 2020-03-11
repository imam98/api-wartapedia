package news

type ListerService interface {
	GetNews(url string) ([]*News, error)
}

type QueryService interface {
	Query(query string) ([]*News, error)
}

type CrawlerService interface {
	Crawl(url string) error
}
