package crawling

import "github.com/imam98/api-wartapedia/pkg/news"

type NewsFetcher interface {
	Fetch(url string) ([]*news.News, error)
}

type crawling struct {
	repo news.Repository
	nf   NewsFetcher
}

func NewCrawler(repo news.Repository, fetcher NewsFetcher) *crawling {
	return &crawling{
		repo: repo,
		nf:   fetcher,
	}
}

func (c *crawling) Crawl(url string) error {
	data, err := c.nf.Fetch(url)
	if err != nil {
		return err
	}

	for _, val := range data {
		if err := c.repo.Store(val); err != nil {
			return err
		}
	}

	return nil
}
