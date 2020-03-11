package listing

import "github.com/imam98/api-wartapedia/pkg/news"

type NewsFetcher interface {
	Fetch(url string) ([]*news.News, error)
}

type listing struct{
	nf NewsFetcher
}

func NewService(fetcher NewsFetcher) news.ListerService {
	return &listing{
		nf: fetcher,
	}
}

func (l *listing) GetNews(url string) ([]*news.News, error) {
	data, err := l.nf.Fetch(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
