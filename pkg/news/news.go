package news

type News struct {
	ID           string `xml:",omitempty"`
	Source       string `xml:",omitempty"`
	Title        string `xml:"title,chardata"`
	MediaContent string `xml:",omitempty"`
	Url          string `xml:"link"`
	Description  string `xml:"description,chardata"`
	PubDate      string `xml:"pubDate"`
}
