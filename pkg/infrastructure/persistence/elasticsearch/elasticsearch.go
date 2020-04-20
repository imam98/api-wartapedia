package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/imam98/api-wartapedia/pkg/news"
)

type repository struct {
	client    *elasticsearch.Client
	indexName string
}

type Config struct {
	Client    *elasticsearch.Client
	IndexName string
}

func NewRepository(config Config) *repository {
	return &repository{client: config.Client, indexName: config.IndexName}
}

func (r *repository) Store(news news.News) error {
	return nil
}

func (r *repository) FindByQuery(query string) ([]news.News, error) {
	return nil, nil
}

func (r *repository) Find(key string) (news.News, error) {
	return news.News{}, nil
}
