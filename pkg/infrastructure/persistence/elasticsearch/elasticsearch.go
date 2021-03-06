package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/imam98/api-wartapedia/pkg/domain"
	"github.com/imam98/api-wartapedia/pkg/domain/entity"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type repository struct {
	client    *elasticsearch.Client
	indexName string
	timeLoc   *time.Location
}

type Config struct {
	Client    *elasticsearch.Client
	IndexName string
	TimeLoc   *time.Location
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
	"sort": { "_score": "desc", "pub_date": "desc" }`

func NewRepository(config Config) *repository {
	if config.TimeLoc == nil {
		config.TimeLoc, _ = time.LoadLocation("Asia/Jakarta")
	}

	return &repository{
		client:    config.Client,
		indexName: config.IndexName,
		timeLoc:   config.TimeLoc,
	}
}

func (r *repository) Store(val entity.News) error {
	pubdate := time.Unix(val.PubDate, 0).In(r.timeLoc)
	expirationDate := time.Now().In(r.timeLoc).AddDate(0, 0, -2)
	if pubdate.YearDay() <= expirationDate.YearDay() && pubdate.Year() <= expirationDate.Year() {
		return domain.ErrItemExpired
	}

	indexName := fmt.Sprintf("%s-%s", r.indexName, pubdate.Format("02-01-2006"))
	resp, err := r.client.Exists(indexName, val.ID)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return domain.ErrItemDuplicate
	}

	payload, err := json.Marshal(val)
	if err != nil {
		return err
	}

	resp, err = r.client.Create(indexName, val.ID, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

func (r *repository) DeleteExpiredIndex() error {
	estimatedDate := time.Now().In(r.timeLoc).AddDate(0, 0, -2)
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
		errType, _ := jsonparser.GetString(data, "error", "type")
		if errType == "index_not_found_exception" {
			return domain.ErrItemNotFound
		}
		reason, _ := jsonparser.GetString(data, "error", "reason")
		return fmt.Errorf("error: %s", reason)
	}

	return nil
}

func (r *repository) FindByQuery(query string, limit int) ([]entity.News, error) {
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
		return []entity.News{}, nil
	}

	var result []entity.News
	var innerErr error = nil
	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		var n entity.News

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

func (r *repository) Find(key string) (entity.News, error) {
	resp, err := r.client.Get(r.indexName, key, r.client.Get.WithPretty())
	if err != nil {
		return entity.News{}, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.News{}, err
	}

	isFound, err := jsonparser.GetBoolean(data, "found")
	if err != nil {
		return entity.News{}, err
	}
	if !isFound {
		return entity.News{}, domain.ErrItemNotFound
	}

	srcObj, _, _, err := jsonparser.Get(data, "_source")
	if err != nil {
		return entity.News{}, nil
	}

	var newsData entity.News
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
