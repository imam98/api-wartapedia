package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/imam98/api-wartapedia/pkg/news"
	"io/ioutil"
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
	resp, err := esapi.GetRequest{
		Index:      r.indexName,
		DocumentID: key,
		Pretty:     true,
	}.Do(context.Background(), r.client)
	if err != nil {
		return news.News{}, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return news.News{}, err
	}

	isFound, err := jsonparser.GetBoolean(data, "found")
	if err != nil {
		return news.News{}, err
	}
	if !isFound {
		return news.News{}, news.ErrItemNotFound
	}

	srcObj, _, _, err := jsonparser.Get(data, "_source")
	if err != nil {
		return news.News{}, nil
	}

	var newsData news.News
	json.Unmarshal(srcObj, &newsData)

	return newsData, nil
}
