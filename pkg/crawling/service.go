package crawling

import (
	"errors"
	"github.com/imam98/api-wartapedia/pkg/news"
)

type NewsFetcher interface {
	Fetch(url string) ([]news.News, error)
}

type Repository interface {
	Find(key string) (news.News, error)
	Store(news news.News) error
}

var ErrSourceNotFound = errors.New("the source flag is not registered in the sourcelist")

type crawling struct {
	repo Repository
	nf   NewsFetcher
}

func NewCrawler(repo Repository, fetcher NewsFetcher) *crawling {
	return &crawling{
		repo: repo,
		nf:   fetcher,
	}
}

func (c *crawling) Crawl(flags news.SourceFlag) error {
	url, ok := news.Sources[flags]
	if !ok {
		return ErrSourceNotFound
	}

	data, err := c.nf.Fetch(url)
	if err != nil {
		return err
	}

	for _, val := range data {
		if _, err := c.repo.Find(val.ID); err == nil {
			break
		}

		val.Source = genSourceFromFlags(flags)
		if err := c.repo.Store(val); err != nil {
			return err
		}
	}

	return nil
}

func genSourceFromFlags(flags news.SourceFlag) string {
	switch flags.SourceOnly() {
	case news.ANTARANEWS:
		return "AntaraNews"
	case news.BBC:
		return "BBC"
	case news.DETIK:
		return "Detik"
	case news.OKEZONE:
		return "Okezone"
	case news.REPUBLIKA:
		return "Republika"
	default:
		return "-"
	}
}
