package news_fetcher

import (
	"encoding/xml"
	"github.com/imam98/api-wartapedia/pkg/news"
	"net/http"
	"regexp"
)

type fetcher struct{}

type newsResults struct {
	XMLName xml.Name    `xml:"rss"`
	N       []news.News `xml:"channel>item"`
}

func NewFetcher() *fetcher {
	return &fetcher{}
}

func (f *fetcher) Fetch(url string) ([]news.News, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data newsResults
	if err := xml.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	matched, err := regexp.MatchString(`(detik|antaranews)`, url)
	if err != nil {
		return nil, err
	}
	if matched {
		re, err := regexp.Compile(`<img src=\"([a-z0-9A-Z:\/\-._]+).*/?>`)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(data.N); i++ {
			subs := re.FindStringSubmatch(data.N[i].Description)
			data.N[i].MediaContent.Src = subs[1]
			data.N[i].Description = re.ReplaceAllString(data.N[i].Description, "")
		}
	}

	return data.N, nil
}
