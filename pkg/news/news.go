package news

import "encoding/xml"

type News struct {
	ID           string      `xml:"-" json:"id,omitempty"`
	Source       string      `xml:"-" json:"source,omitempty"`
	Title        string      `xml:"title" json:"title"`
	Enclosure    Media       `xml:"enclosure,omitempty" json:"enclosure,omitempty"`
	MediaContent Media       `xml:"content,omitempty" json:"media_content,omitempty"`
	Url          string      `xml:"link" json:"url"`
	Description  Description `xml:"description" json:"description"`
	PubDate      string      `xml:"pubDate" json:"pub_date"`
}

type Description struct {
	XMLName xml.Name `xml:"description" json:"-"`
	Text    string   `xml:",cdata" json:"text"`
}

type Media struct {
	Src string `xml:"url,attr" json:"src,omitempty"`
}
