package crawling

import (
	"errors"
	"github.com/imam98/api-wartapedia/pkg/news"
	"regexp"
	"strings"
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

		source, prefix := parseSourceFromFlags(flags)
		val.Source = source
		val.ID = genDocID(prefix, val.Url)
		if err := c.repo.Store(val); err != nil {
			return err
		}
	}

	return nil
}

func parseSourceFromFlags(flags news.SourceFlag) (source, prefix string) {
	switch flags.SourceOnly() {
	case news.ANTARANEWS:
		source = "AntaraNews"
		prefix = "atn"
	case news.BBC:
		source = "BBC"
		prefix = "bbc"
	case news.DETIK:
		source = "Detik"
		prefix = "dtk"
	case news.OKEZONE:
		source = "Okezone"
		prefix = "okz"
	case news.REPUBLIKA:
		source = "Republika"
		prefix = "rpb"
	default:
		source = "-"
		prefix = "-"
	}

	return
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
