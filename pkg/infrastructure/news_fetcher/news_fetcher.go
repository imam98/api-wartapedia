package news_fetcher

import (
	"encoding/xml"
	"github.com/imam98/api-wartapedia/pkg/news"
	"html"
	"net/http"
	"regexp"
	"time"
)

type fetcher struct{}

type fetchResult struct {
	XMLName xml.Name     `xml:"rss"`
	N       []newsResult `xml:"channel>item"`
}

type newsResult struct {
	Title        string `xml:"title"`
	MediaContent media  `xml:"content,omitempty"`
	Url          string `xml:"link"`
	Description  string `xml:"description"`
	PubDate      string `xml:"pubDate"`
}

type media struct {
	Src string `xml:"url,attr"`
}

func NewFetcher() *fetcher {
	return &fetcher{}
}

func (f *fetcher) Fetch(url string) ([]news.News, error) {
	var results []news.News
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

	var data fetchResult
	if err := xml.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	mustParseMediaContent, err := regexp.MatchString(`(detik|antaranews)`, url)
	if err != nil {
		return nil, err
	}

	namedTimezone, err := regexp.MatchString(`bbc`, url)
	if err != nil {
		return nil, err
	}

	for _, val := range data.N {
		err := cleanData(&val, mustParseMediaContent)
		if err != nil {
			return nil, err
		}

		unixTime, err := parseUnixTime(val.PubDate, namedTimezone)
		if err != nil {
			return nil, err
		}

		n := news.News{
			Title:        val.Title,
			MediaContent: val.MediaContent.Src,
			Url:          val.Url,
			Description:  val.Description,
			PubDate:      unixTime,
		}

		results = append(results, n)
	}

	return results, nil
}

func cleanData(val *newsResult, mustParseMediaContent bool) error {
	if mustParseMediaContent {
		err := parseMediaContent(val)
		if err != nil {
			return err
		}
	}

	val.Description = html.UnescapeString(val.Description)

	return nil
}

func parseMediaContent(val *newsResult) error {
	re, err := regexp.Compile(`<img src=\"([a-z0-9A-Z:\/\-._]+).*/?>`)
	if err != nil {
		return err
	}

	subs := re.FindStringSubmatch(val.Description)
	val.MediaContent.Src = subs[1]
	val.Description = re.ReplaceAllString(val.Description, "")

	return nil
}

func parseUnixTime(timeValue string, namedTimezone bool) (int64, error) {
	var layout string
	if namedTimezone {
		layout = time.RFC1123
	} else {
		layout = time.RFC1123Z
	}

	parsedTime, err := time.Parse(layout, timeValue)
	if err != nil {
		return 0, err
	}

	return parsedTime.Unix(), nil
}
