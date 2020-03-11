package querying

import "github.com/imam98/api-wartapedia/pkg/news"

type querying struct {
	repo news.Repository
}

func NewService(repo news.Repository) *querying {
	return &querying{
		repo: repo,
	}
}

func (q *querying) Query(query string) ([]*news.News, error) {
	data, err := q.repo.FindByQuery(query)
	if err != nil {
		return nil, err
	}

	return data, nil
}
