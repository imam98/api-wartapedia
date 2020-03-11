package news_fetcher

import (
	"encoding/json"
	"github.com/imam98/api-wartapedia/pkg/news"
	"net/http"
)

type fetcher struct{}

type newsResults struct {
	n []*news.News
}

func NewFetcher() *fetcher {
	return &fetcher{}
}

func (f *fetcher) Fetch(url string) ([]*news.News, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data newsResults
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.n, nil
}
