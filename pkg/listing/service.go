package listing

import (
	"errors"
	"github.com/imam98/api-wartapedia/pkg/news"
	"sort"
)

type NewsFetcher interface {
	Fetch(url string) ([]news.News, error)
}

type listing struct {
	nf NewsFetcher
}

var ErrSourceNotFound = errors.New("the source flag is not registered in the sourcelist")
var ErrInvalidCategoryFlag = errors.New("category flag is invalid")

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

func (l *listing) GetPublishersFromCategory(catFlag news.SourceFlag) ([]string, error) {
	if catFlag.CategoryOnly() != catFlag {
		return nil, ErrInvalidCategoryFlag
	}

	var categories []string
	for idx, _ := range news.Sources {
		if catFlag^idx.CategoryOnly() == 0 {
			cat := idx.SourceString()
			categories = append(categories, cat)
		}
	}

	sort.Strings(categories)
	return categories, nil
}


