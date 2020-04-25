package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/imam98/api-wartapedia/pkg/news"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type repository struct {
	client    *elasticsearch.Client
	indexName string
}

type Config struct {
	Client    *elasticsearch.Client
	IndexName string
}

const searchAll = `
	"query": { "match_all": {} },
	"size": %d,
	"sort": { "pub_date": "desc" }`

const searchQuery = `
	"query": {
		"multi_match": {
			"query": %q,
			"fields": ["title", "description"],
			"operator": "or"
		}
	},
	"size": %d,
	"sort": { "_score": "asc", "pub_date": "desc" }`

func NewRepository(config Config) *repository {
	return &repository{client: config.Client, indexName: config.IndexName}
}

func (r *repository) Store(val news.News) error {
	pubdate := time.Unix(val.PubDate, 0)
	expirationDate := time.Now().AddDate(0, 0, -2)
	if pubdate.Equal(expirationDate) || pubdate.Before(expirationDate) {
		return news.ErrItemExpired
	}

	indexName := fmt.Sprintf("%s-%s", r.indexName, pubdate.Format("02-01-2006"))
	resp, err := r.client.Exists(indexName, val.ID)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return news.ErrItemDuplicate
	}

	payload, err := json.Marshal(val)
	if err != nil {
		return err
	}

	if _, err := r.client.Create(r.indexName, val.ID, bytes.NewReader(payload)); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteExpiredIndex() error {
	estimatedDate := time.Now().AddDate(0, 0, -2)
	indexName := fmt.Sprintf("%s-%s", r.indexName, estimatedDate.Format("02-01-2006"))
	resp, err := r.client.Indices.Delete([]string{indexName})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.IsError() {
		reason, _ := jsonparser.GetString(data, "error", "reason")
		return fmt.Errorf("error: %s", reason)
	}

	return nil
}

func (r *repository) FindByQuery(query string, limit int) ([]news.News, error) {
	resp, err := r.client.Search(
		r.client.Search.WithBody(matchQueryBuilder(query, limit)),
		r.client.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	total, err := jsonparser.GetInt(data, "hits", "total", "value")
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return []news.News{}, nil
	}

	var result []news.News
	var innerErr error = nil
	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		var n news.News

		srcObj, _, _, err := jsonparser.Get(value, "_source")
		if err != nil {
			innerErr = err
			return
		}
		json.Unmarshal(srcObj, &n)
		result = append(result, n)
	}, "hits", "hits")
	if err != nil {
		return nil, err
	}
	if innerErr != nil {
		return nil, innerErr
	}

	return result, nil
}

func (r *repository) Find(key string) (news.News, error) {
	resp, err := r.client.Get(r.indexName, key, r.client.Get.WithPretty())
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

func matchQueryBuilder(query string, limit int) io.Reader {
	var sb strings.Builder

	sb.WriteString("{\n")
	if query == "" {
		sb.WriteString(fmt.Sprintf(searchAll, limit))
	} else {
		sb.WriteString(fmt.Sprintf(searchQuery, query, limit))
	}
	sb.WriteString("\n}")

	return strings.NewReader(sb.String())
}
