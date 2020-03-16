package news

import "encoding/xml"

type News struct {
	ID           string      `xml:"-"`
	Source       string      `xml:"-"`
	Title        string      `xml:"title"`
	Enclosure    Media       `xml:"enclosure,omitempty"`
	MediaContent Media       `xml:"content,omitempty"`
	Url          string      `xml:"link"`
	Description  Description `xml:"description"`
	PubDate      string      `xml:"pubDate"`
}

type Description struct {
	XMLName xml.Name `xml:"description"`
	Text    string   `xml:",cdata"`
}

type Media struct {
	Src string `xml:"url,attr"`
}
