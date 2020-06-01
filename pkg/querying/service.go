package querying

import (
	"github.com/imam98/api-wartapedia/pkg/domain/entity"
)

type Repository interface {
	FindByQuery(query string, limit int) ([]entity.News, error)
}

type querying struct {
	repo Repository
}

func NewService(repo Repository) *querying {
	return &querying{
		repo: repo,
	}
}

func (q *querying) Query(query string, limit int) ([]entity.News, error) {
	data, err := q.repo.FindByQuery(query, limit)
	if err != nil {
		return nil, err
	}

	return data, nil
}
