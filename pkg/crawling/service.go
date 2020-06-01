package crawling

import (
	"github.com/imam98/api-wartapedia/pkg/domain"
	"github.com/imam98/api-wartapedia/pkg/domain/entity"
	"regexp"
	"strings"
)

type NewsFetcher interface {
	Fetch(url string) ([]entity.News, error)
}

type Repository interface {
	Find(key string) (entity.News, error)
	Store(news entity.News) error
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

func (c *crawling) Crawl(flags domain.RepoFlag) error {
	url, ok := domain.Sources[flags]
	if !ok {
		return domain.ErrSourceNotFound
	}

	data, err := c.nf.Fetch(url)
	if err != nil {
		return err
	}

	for _, val := range data {
		val.Source = flags.SourceString()
		prefix := parsePrefixFromFlags(flags)
		val.ID = genDocID(prefix, val.Url)
		if err := c.repo.Store(val); err != nil {
			if err == domain.ErrItemDuplicate {
				break
			} else if err == domain.ErrItemExpired {
				break
			} else {
				return err
			}
		}
	}

	return nil
}

func parsePrefixFromFlags(flags domain.RepoFlag) string {
	switch flags.SourceOnly() {
	case domain.ANTARANEWS:
		return "atn"
	case domain.BBC:
		return "bbc"
	case domain.DETIK:
		return "dtk"
	case domain.OKEZONE:
		return "okz"
	case domain.REPUBLIKA:
		return "rpb"
	default:
		return "-"
	}
}

func genDocID(prefix, url string) string {
	sb := strings.Builder{}
	sb.WriteString(prefix)
	sb.WriteString("::")

	re, _ := regexp.Compile("(http://|https://)")
	url = re.ReplaceAllString(url, "")
	segment := strings.Split(url, "/")
	id := segment[len(segment)-1]
	sb.WriteString(id)

	return sb.String()
}
