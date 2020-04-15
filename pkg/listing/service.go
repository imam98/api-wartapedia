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

var ErrInvalidCategoryFlag = errors.New("category flag is invalid")

func NewService(fetcher NewsFetcher) news.ListerService {
	return &listing{
		nf: fetcher,
	}
}

func (l *listing) GetNews(flags news.RepoFlag) ([]news.News, error) {
	url, ok := news.Sources[flags]
	if !ok {
		return nil, news.ErrSourceNotFound
	}

	data, err := l.nf.Fetch(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *listing) GetSourcesFromCategory(catFlag news.RepoFlag) ([]string, error) {
	if catFlag.CategoryOnly() != catFlag {
		return nil, ErrInvalidCategoryFlag
	}

	var sources []string
	for i := 1; i <= 5; i++ {
		flags := catFlag | news.RepoFlag(i)
		if flags.Validate() {
			src := flags.SourceString()
			sources = append(sources, src)
		}
	}

	sort.Strings(sources)
	return sources, nil
}
