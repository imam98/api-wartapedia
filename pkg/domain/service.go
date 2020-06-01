package domain

import "github.com/imam98/api-wartapedia/pkg/domain/entity"

type ListerService interface {
	GetNews(flags RepoFlag) ([]entity.News, error)
	GetSourcesFromCategory(flags RepoFlag) ([]string, error)
}

type QueryService interface {
	Query(query string, limit int) ([]entity.News, error)
}

type CrawlerService interface {
	Crawl(flags RepoFlag) error
}

type DeleterDaemonService interface {
	Start() error
}
