package crawling

import (
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

func (c *crawling) Crawl(flags news.RepoFlag) error {
	url, ok := news.Sources[flags]
	if !ok {
		return news.ErrSourceNotFound
	}

	data, err := c.nf.Fetch(url)
	if err != nil {
		return err
	}

	for _, val := range data {
		if _, err := c.repo.Find(val.ID); err == nil {
			break
		}

		val.Source = flags.SourceString()
		prefix := parsePrefixFromFlags(flags)
		val.ID = genDocID(prefix, val.Url)
		if err := c.repo.Store(val); err != nil {
			return err
		}
	}

	return nil
}

func parsePrefixFromFlags(flags news.RepoFlag) string {
	switch flags.SourceOnly() {
	case news.ANTARANEWS:
		return "atn"
	case news.BBC:
		return "bbc"
	case news.DETIK:
		return "dtk"
	case news.OKEZONE:
		return "okz"
	case news.REPUBLIKA:
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
