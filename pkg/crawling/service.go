package crawling

import "github.com/imam98/api-wartapedia/pkg/news"

type NewsFetcher interface {
	Fetch(url string) ([]news.News, error)
}

type Repository interface {
	Find(key string) (news.News, error)
	Store(news news.News) error
}

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

func (c *crawling) Crawl(url string) error {
	data, err := c.nf.Fetch(url)
	if err != nil {
		return err
	}

	for _, val := range data {
		if _, err := c.repo.Find(val.ID); err == nil {
			break
		}

		if err := c.repo.Store(val); err != nil {
			return err
		}
	}

	return nil
}
