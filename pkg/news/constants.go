package news

import "errors"

const (
	ANTARANEWS = iota + 1
	BBC
	DETIK
	OKEZONE
	REPUBLIKA
)

const (
	CAT_NASIONAL = 16
	CAT_DUNIA    = 32
	CAT_TEKNO    = 48
)

type RepoFlag byte

var Sources = map[RepoFlag]string{
	CAT_NASIONAL | ANTARANEWS: "https://www.antaranews.com/rss/terkini",
	CAT_NASIONAL | BBC:        "http://feeds.bbci.co.uk/indonesia/rss.xml",
	CAT_NASIONAL | DETIK:      "http://rss.detik.com/index.php/detiknews",
	CAT_NASIONAL | OKEZONE:    "http://sindikasi.okezone.com/index.php/rss/1/RSS2.0",
	CAT_NASIONAL | REPUBLIKA:  "https://www.republika.co.id/rss/nasional/",
	CAT_DUNIA | ANTARANEWS:    "https://www.antaranews.com/rss/dunia",
	CAT_DUNIA | BBC:           "http://feeds.bbci.co.uk/indonesia/dunia/rss.xml",
	CAT_DUNIA | DETIK:         "http://rss.detik.com/index.php/detikcom_internasional",
	CAT_DUNIA | REPUBLIKA:     "http://www.republika.co.id/rss/internasional/",
	CAT_TEKNO | ANTARANEWS:    "https://www.antaranews.com/rss/tekno",
	CAT_TEKNO | OKEZONE:       "http://sindikasi.okezone.com/index.php/rss/16/RSS2.0",
	CAT_TEKNO | REPUBLIKA:     "https://www.republika.co.id/rss/leisure/oto-tek/",
}

var (
	ErrItemNotFound  = errors.New("Error item not found")
	ErrItemDuplicate = errors.New("Error item already exists")
	ErrSourceNotFound = errors.New("the source flag is not registered in the sourcelist")
)

func (r RepoFlag) SourceOnly() RepoFlag {
	return r & RepoFlag(15)
}

func (r RepoFlag) CategoryOnly() RepoFlag {
	return r & RepoFlag(240)
}

func (r RepoFlag) SourceString() string {
	switch r.SourceOnly() {
	case ANTARANEWS:
		return "AntaraNews"
	case BBC:
		return "BBC"
	case DETIK:
		return "Detik"
	case OKEZONE:
		return "Okezone"
	case REPUBLIKA:
		return "Republika"
	default:
		return "-"
	}
}
