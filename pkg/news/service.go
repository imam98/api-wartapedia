package news

type ListerService interface {
	GetNews(flags SourceFlag) ([]News, error)
}

type QueryService interface {
	Query(query string) ([]News, error)
}

type CrawlerService interface {
	Crawl(flags SourceFlag) error
}
