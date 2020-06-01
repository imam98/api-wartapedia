package entity

type News struct {
	ID           string `json:"id,omitempty"`
	Source       string `json:"source,omitempty"`
	Title        string `json:"title"`
	MediaContent string `json:"media_content,omitempty"`
	Url          string `json:"url"`
	Description  string `json:"description"`
	PubDate      int64  `json:"pub_date"`
}
