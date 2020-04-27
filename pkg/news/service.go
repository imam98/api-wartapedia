package news

type ListerService interface {
	GetNews(flags RepoFlag) ([]News, error)
	GetSourcesFromCategory(flags RepoFlag) ([]string, error)
}

type QueryService interface {
	Query(query string, limit int) ([]News, error)
}

type CrawlerService interface {
	Crawl(flags RepoFlag) error
}
