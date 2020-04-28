package main

import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/imam98/api-wartapedia/pkg/deleting"
	"github.com/imam98/api-wartapedia/pkg/infrastructure/persistence/elasticsearch"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func main() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC822Z,
	}
	logger := zerolog.New(output).With().Timestamp().Logger()

	client, err := es.NewDefaultClient()
	if err != nil {
		logger.Error().Err(err).Msg("Elasticsearch error")
	}

	repo := elasticsearch.NewRepository(elasticsearch.Config{
		Client:    client,
		IndexName: "wartapedia",
	})
	service := deleting.NewService(repo)
	if err := service.Start(); err != nil {
		logger.Error().Err(err).Msg("Unexpected error")
	}
}
