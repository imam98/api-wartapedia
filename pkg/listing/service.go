package listing

import (
	"errors"
	"github.com/imam98/api-wartapedia/pkg/domain"
	"github.com/imam98/api-wartapedia/pkg/domain/entity"
	"sort"
)

type NewsFetcher interface {
	Fetch(url string) ([]entity.News, error)
}

type listing struct {
	nf NewsFetcher
}

var ErrInvalidCategoryFlag = errors.New("category flag is invalid")

func NewService(fetcher NewsFetcher) domain.ListerService {
	return &listing{
		nf: fetcher,
	}
}

func (l *listing) GetNews(flags domain.RepoFlag) ([]entity.News, error) {
	url, ok := domain.Sources[flags]
	if !ok {
		return nil, domain.ErrSourceNotFound
	}

	data, err := l.nf.Fetch(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (l *listing) GetSourcesFromCategory(catFlag domain.RepoFlag) ([]string, error) {
	if catFlag.CategoryOnly() != catFlag {
		return nil, ErrInvalidCategoryFlag
	}

	var sources []string
	for i := 1; i <= 5; i++ {
		flags := catFlag | domain.RepoFlag(i)
		if flags.Validate() {
			src := flags.SourceString()
			sources = append(sources, src)
		}
	}

	sort.Strings(sources)
	return sources, nil
}
