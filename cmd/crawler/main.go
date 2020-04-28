package main

import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/imam98/api-wartapedia/pkg/crawling"
	"github.com/imam98/api-wartapedia/pkg/infrastructure/news_fetcher"
	"github.com/imam98/api-wartapedia/pkg/infrastructure/persistence/elasticsearch"
	"github.com/imam98/api-wartapedia/pkg/news"
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
	fetcher := news_fetcher.NewFetcher()
	crawler := crawling.NewCrawler(repo, fetcher)

	for {
		for flags, _ := range news.Sources {
			go func(f news.RepoFlag) {
				logger.Info().
					Str("source", f.SourceString()).
					Str("category", f.CategoryString()).
					Msg("Crawling has been started")
				if err := crawler.Crawl(f); err != nil {
					logger.Error().
						Err(err).
						Str("source", f.SourceString()).
						Str("category", f.CategoryString()).
						Msg("Crawler error")
				}
				logger.Info().
					Str("source", f.SourceString()).
					Str("category", f.CategoryString()).
					Msg("Crawling has been finished")
			}(flags)
		}

		time.Sleep(60 * time.Minute)
	}
}
