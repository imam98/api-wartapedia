package news

import "encoding/xml"

type News struct {
	ID           string      `xml:"-"`
	Source       string      `xml:"-"`
	Title        string      `xml:"title"`
	MediaContent Media       `xml:"enclosure,omitempty"`
	Url          string      `xml:"link"`
	Description  Description `xml:"description"`
	PubDate      string      `xml:"pubDate"`
}

type Description struct {
	XMLName xml.Name `xml:"description"`
	Text    string   `xml:",cdata"`
}

type Media struct {
	XMLName xml.Name `xml:"enclosure"`
	Src     string   `xml:"url,attr"`
}
