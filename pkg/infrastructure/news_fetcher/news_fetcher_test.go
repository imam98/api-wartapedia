package news_fetcher

import (
	"fmt"
	"github.com/imam98/api-wartapedia/pkg/news"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fakeAntaraServer(w http.ResponseWriter, r *http.Request) {
	responseXML := `
	<?xml version="1.0" encoding="UTF-8"?>
	<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/" xmlns:slash="http://purl.org/rss/1.0/modules/slash/" version="2.0">
	  <channel>
		<title>ANTARA News - Berita Terkini</title>
		<description>News And Service</description>
		<link>https://www.antaranews.com</link>
		<language>id</language>
		<copyright>2020 ANTARA News</copyright>
		<lastBuildDate>Sun, 15 Mar 2020 15:54:01 +0700</lastBuildDate>
		<atom:link href="https://www.antaranews.com/rss/terkini.xml" rel="self" type="application/rss+xml"/>
		<item>
		  <title>Dummy Title</title>
		  <link>https://www.antaranews.com/berita/1357722/surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota</link>
		  <pubDate>Sun, 15 Mar 2020 15:51:04 +0700</pubDate>
		  <description><![CDATA[<img src="dummy.jpg">Dummy description]]></description>
		  <guid isPermaLink="false">https://www.antaranews.com/berita/1357722/surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota</guid>
		</item>
		<item>
		  <title>Dummy Title</title>
		  <link>https://www.antaranews.com/foto/1357714/pencegahan-wabah-covid-19-di-kalimantan-tengah</link>
		  <pubDate>Sun, 15 Mar 2020 15:51:04 +0700</pubDate>
		  <description><![CDATA[<img src="dummy.jpg">Dummy description]]></description>
		  <guid isPermaLink="false">https://www.antaranews.com/foto/1357714/pencegahan-wabah-covid-19-di-kalimantan-tengah</guid>
		</item>
		<item>
		  <title>Dummy Title</title>
		  <link>https://www.antaranews.com/video/1357690/presiden-imbau-masyarakat-bekerja-belajar-dan-beribadah-di-rumah</link>
		  <pubDate>Sun, 15 Mar 2020 15:51:04 +0700</pubDate>
		  <description><![CDATA[<img src="dummy.jpg">Dummy description]]></description>
		  <guid isPermaLink="false">https://www.antaranews.com/video/1357690/presiden-imbau-masyarakat-bekerja-belajar-dan-beribadah-di-rumah</guid>
		</item>
	  </channel>
	</rss>
	`

	w.Header().Set("content-type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, responseXML)
}

func fakeBBCServer(w http.ResponseWriter, r *http.Request) {
	responseXML := `
	<?xml version="1.0" encoding="UTF-8"?>
	<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
		<channel>
			<title><![CDATA[BBC News Indonesia - Berita]]></title>
			<description><![CDATA[BBC News Indonesia - Berita]]></description>
			<link>http://www.bbcindonesia.com</link>
			<image>
				<url>http://www.bbc.co.uk/indonesia/images/gel/rss_logo.gif</url>
				<title>BBC News Indonesia - Berita</title>
				<link>http://www.bbcindonesia.com</link>
			</image>
			<generator>RSS for Node</generator>
			<lastBuildDate>Thu, 12 Mar 2020 05:08:20 GMT</lastBuildDate>
			<copyright><![CDATA[Hak cipta British Broadcasting Corporation ]]></copyright>
			<language><![CDATA[id]]></language>
			<managingEditor><![CDATA[bbcindonesia@bbc.co.uk]]></managingEditor>
			<ttl>15</ttl>
			<item>
				<title><![CDATA[Virus corona: Mengapa Indonesia 'tidak terbuka', sementara negara lain bersikap 'transparan'?]]></title>
				<description><![CDATA[Dummy description]]></description>
				<link>http://www.bbc.com/indonesia/indonesia-51842758</link>
				<guid isPermaLink="true">http://www.bbc.com/indonesia/indonesia-51842758</guid>
				<pubDate>Thu, 12 Mar 2020 04:40:23 GMT</pubDate>
			</item>
			<item>
				<title><![CDATA[Sejarah bulu tangkis di Olimpiade: Mengapa Indonesia sulit lahirkan Susy Susanti generasi baru?]]></title>
				<description><![CDATA[Dummy description]]></description>
				<link>http://www.bbc.com/indonesia/olahraga-51662063</link>
				<guid isPermaLink="true">http://www.bbc.com/indonesia/olahraga-51662063</guid>
				<pubDate>Thu, 12 Mar 2020 02:42:21 GMT</pubDate>
			</item>
			<item>
				<title><![CDATA[Virus corona: Karyawan apresiasi pembebasan pajak penghasilan, ekonom sebut 'perlu stimulus fiskal dan moneter' atasi perlambatan ekonomi]]></title>
				<description><![CDATA[Dummy description]]></description>
				<link>http://www.bbc.com/indonesia/indonesia-51830029</link>
				<guid isPermaLink="true">http://www.bbc.com/indonesia/indonesia-51830029</guid>
				<pubDate>Thu, 12 Mar 2020 01:21:54 GMT</pubDate>
			</item>
		</channel>
	</rss>
	`
	w.Header().Set("content-type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, responseXML)
}

func TestBBCNewsFetcher(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fakeBBCServer))
	defer server.Close()

	expectedResult := []news.News{
		news.News{
			Title:       "Virus corona: Mengapa Indonesia 'tidak terbuka', sementara negara lain bersikap 'transparan'?",
			Url:         "http://www.bbc.com/indonesia/indonesia-51842758",
			Description: news.Description{Text: "Dummy description"},
			PubDate:     "Thu, 12 Mar 2020 04:40:23 GMT",
		},
		news.News{
			Title:       "Sejarah bulu tangkis di Olimpiade: Mengapa Indonesia sulit lahirkan Susy Susanti generasi baru?",
			Url:         "http://www.bbc.com/indonesia/olahraga-51662063",
			Description: news.Description{Text: "Dummy description"},
			PubDate:     "Thu, 12 Mar 2020 02:42:21 GMT",
		},
		news.News{
			Title:       "Virus corona: Karyawan apresiasi pembebasan pajak penghasilan, ekonom sebut 'perlu stimulus fiskal dan moneter' atasi perlambatan ekonomi",
			Url:         "http://www.bbc.com/indonesia/indonesia-51830029",
			Description: news.Description{Text: "Dummy description"},
			PubDate:     "Thu, 12 Mar 2020 01:21:54 GMT",
		},
	}

	fetcher := NewFetcher()
	data, err := fetcher.Fetch(server.URL)
	if err != nil {
		t.Fatalf("Error occured: %q", err)
	}

	assertLength(t, expectedResult, data)
	assertElements(t, expectedResult, data)
}

func TestAntaraNewsFetcher(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fakeAntaraServer))
	defer server.Close()

	expectedResult := []news.News{
		news.News{
			Title:       "Dummy Title",
			Url:         "https://www.antaranews.com/berita/1357722/surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota",
			Description: news.Description{Text: "<img src=\"dummy.jpg\">Dummy description"},
			PubDate:     "Sun, 15 Mar 2020 15:51:04 +0700",
		},
		news.News{
			Title:       "Dummy Title",
			Url:         "https://www.antaranews.com/foto/1357714/pencegahan-wabah-covid-19-di-kalimantan-tengah",
			Description: news.Description{Text: "<img src=\"dummy.jpg\">Dummy description"},
			PubDate:     "Sun, 15 Mar 2020 15:51:04 +0700",
		},
		news.News{
			Title:       "Dummy Title",
			Url:         "https://www.antaranews.com/video/1357690/presiden-imbau-masyarakat-bekerja-belajar-dan-beribadah-di-rumah",
			Description: news.Description{Text: "<img src=\"dummy.jpg\">Dummy description"},
			PubDate:     "Sun, 15 Mar 2020 15:51:04 +0700",
		},
	}

	fetcher := NewFetcher()
	data, err := fetcher.Fetch(server.URL)
	if err != nil {
		t.Fatalf("Error occured: %q", err)
	}

	assertLength(t, expectedResult, data)
	assertElements(t, expectedResult, data)
}

func assertLength(t *testing.T, expected []news.News, got []news.News) {
	t.Helper()

	if len(expected) != len(got) {
		t.Fatalf("Size of slice doesn't match!\nExpected: %v\nGot: %v\n", expected, got)
	}
}

func assertElements(t *testing.T, expected []news.News, got []news.News) {
	for index, val := range expected {
		if val.Title != got[index].Title {
			t.Errorf("Struct value doesn't match!\nExpected: %v\nGot: %v\n", val.Title, got[index].Title)
		}

		if val.Description.Text != got[index].Description.Text {
			t.Errorf("Struct value doesn't match!\nExpected: %v\nGot: %v\n", val.Description, got[index].Description)
		}

		if val.PubDate != got[index].PubDate {
			t.Errorf("Struct value doesn't match!\nExpected: %v\nGot: %v\n", val.PubDate, got[index].PubDate)
		}

		if val.Url != got[index].Url {
			t.Errorf("Struct value doesn't match!\nExpected: %v\nGot: %v\n", val.Url, got[index].Url)
		}
	}
}
