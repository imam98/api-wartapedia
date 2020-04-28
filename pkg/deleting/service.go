package deleting

import (
	"github.com/imam98/api-wartapedia/pkg/news"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Repository interface {
	DeleteExpiredIndex() error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo:repo}
}

func (s *service) Start() error {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC822Z,
	}
	logger := zerolog.New(output).With().Timestamp().Logger()

	for {
		logger.Info().Msg("Start delete operation")
		err := s.repo.DeleteExpiredIndex()
		if err != nil {
			if err == news.ErrItemNotFound {
				logger.Error().Msg("Index not found exception")
			} else {
				return err
			}
		}

		logger.Info().Msg("Delete operation finished")
		time.Sleep(getTimeEstimation(loc))
	}
}

func getTimeEstimation(loc *time.Location) time.Duration {
	currTime := time.Now().In(loc)
	nextDay := currTime.AddDate(0, 0, 1)
	rounded := time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, loc)
	estimation := rounded.Sub(currTime)
	return estimation
}
