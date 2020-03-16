package listing

import (
	"errors"
	"github.com/imam98/api-wartapedia/pkg/news"
)

type NewsFetcher interface {
	Fetch(url string) ([]news.News, error)
}

type listing struct {
	nf NewsFetcher
}

var ErrSourceNotFound = errors.New("the source flag is not registered in the sourcelist")

func NewService(fetcher NewsFetcher) news.ListerService {
	return &listing{
		nf: fetcher,
	}
}

func (l *listing) GetNews(flags news.SourceFlag) ([]news.News, error) {
	url, ok := news.Sources[flags]
	if !ok {
		return nil, ErrSourceNotFound
	}

	data, err := l.nf.Fetch(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
