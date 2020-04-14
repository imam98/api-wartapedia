package news

type ListerService interface {
	GetNews(flags RepoFlag) ([]News, error)
	GetSourcesFromCategory(flags RepoFlag) ([]string, error)
}

type QueryService interface {
	Query(query string) ([]News, error)
}

type CrawlerService interface {
	Crawl(flags RepoFlag) error
}
