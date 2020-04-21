package news

type News struct {
	ID           string `xml:"-" json:"id,omitempty"`
	Source       string `xml:"-" json:"source,omitempty"`
	Title        string `xml:"title" json:"title"`
	MediaContent Media  `xml:"content,omitempty" json:"media_content,omitempty"`
	Url          string `xml:"link" json:"url"`
	Description  string `xml:"description" json:"description"`
	PubDate      string `xml:"pubDate" json:"pub_date"`
}

type Media struct {
	Src string `xml:"url,attr" json:"src,omitempty"`
}
