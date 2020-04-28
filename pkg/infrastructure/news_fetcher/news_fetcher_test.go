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
		  <description><![CDATA[<img src="https://dummy.jpg" align="left" border="0">Dummy description]]></description>
		  <guid isPermaLink="false">https://www.antaranews.com/berita/1357722/surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota</guid>
		</item>
		<item>
		  <title>Dummy Title</title>
		  <link>https://www.antaranews.com/foto/1357714/pencegahan-wabah-covid-19-di-kalimantan-tengah</link>
		  <pubDate>Sun, 15 Mar 2020 15:51:04 +0700</pubDate>
		  <description><![CDATA[<img src="https://dummy.jpg" align="left" border="0">Dummy description]]></description>
		  <guid isPermaLink="false">https://www.antaranews.com/foto/1357714/pencegahan-wabah-covid-19-di-kalimantan-tengah</guid>
		</item>
		<item>
		  <title>Dummy Title</title>
		  <link>https://www.antaranews.com/video/1357690/presiden-imbau-masyarakat-bekerja-belajar-dan-beribadah-di-rumah</link>
		  <pubDate>Sun, 15 Mar 2020 15:51:04 +0700</pubDate>
		  <description><![CDATA[<img src="https://dummy.jpg" align="left" border="0">Dummy description]]></description>
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

func fakeDetikServer(w http.ResponseWriter, r *http.Request) {
	responseXML := `
	<?xml version="1.0" encoding="UTF-8"?>
	<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
		<channel>
			<title>news.detik</title>
			<link>http://news.detik.com/</link>
			<description>Detik.com sindikasi</description>
			<image>
				<title>detikNews - Berita</title>
				<link>http://news.detik.com/</link>
				<url>http://rss.detik.com/images/rsslogo_detikcom.gif</url>
			</image>
			<item>
				<title><![CDATA[Dummy Title]]></title>
				<link>https://news.detik.com/read/2020/03/15/182444/4940126/10/belajar-mengajar-tk-sampai-smp-di-kendari-pindah-ke-rumah-imbas-corona</link>
				<guid>https://news.detik.com/read/2020/03/15/182444/4940126/10/belajar-mengajar-tk-sampai-smp-di-kendari-pindah-ke-rumah-imbas-corona</guid>
				<pubDate>Sun, 15 Mar 2020 18:34:46 +0700</pubDate>
				<description>&lt;img src=&quot;https://dummy.jpeg&quot; align=&quot;left&quot; hspace=&quot;7&quot; width=&quot;100&quot; /&gt;<![CDATA["Dummy description"]]></description>
				<enclosure url="https://dummy.png" length="10240" type="image/png" />
			</item>
		</channel>
	</rss>
	`

	w.Header().Set("content-type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, responseXML)
}

func fakeOkezoneServer(w http.ResponseWriter, r *http.Request) {
	responseXML := `
	<?xml version="1.0" encoding="UTF-8"?>
	<rss version="2.0"
		xmlns:content="http://purl.org/rss/1.0/modules/content/"
		xmlns:media="http://search.yahoo.com/mrss/"
		xmlns:atom="http://www.w3.org/2005/Atom">
		<channel>
			<title>Sindikasi news.okezone.com</title>
			<description>Berita-berita Okezone pada kanal News</description>
			<link>https://news.okezone.com</link>
			<lastBuildDate>Mon, 16 Mar 2020 13:53:26 +0700</lastBuildDate>
			<generator>Okezone RSS 2.0 Generator</generator>
			<image>
				<link>https://news.okezone.com</link>
				<title>Sindikasi news.okezone.com</title>
				<url>https://cdn.okezone.com/underwood/revamp/2019/logo/desktop/icon-okz.png</url>
				<description>Berita-berita Okezone pada kanal News</description>
			</image>
			<item>
				<title>Dummy Title</title>
				<link>https://megapolitan.okezone.com/read/2020/03/16/338/2184032/cegah-penyebaran-covid-19-pemkab-bekasi-tiadakan-kegiatan-publik</link>
				<guid>https://megapolitan.okezone.com/read/2020/03/16/338/2184032/cegah-penyebaran-covid-19-pemkab-bekasi-tiadakan-kegiatan-publik</guid>
				<description>Dummy description</description>
				<media:content url="https://dummy.jpg?w=300" 
										type="image/jpg" expression="full" width="300" height="190"></media:content>
				<category>breaking news - Property</category>
				<pubDate>Mon, 16 Mar 2020 13:53:18 +0700</pubDate>
			</item>
			<atom:link href="https://sindikasi.okezone.com/index.php/rss/0/RSS2.0" rel="self" type="application/rss+xml" />
		</channel>
	</rss>
	`

	w.Header().Set("content-type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, responseXML)
}

func fakeRepublikaServer(w http.ResponseWriter, r *http.Request) {
	responseXML := `
	<?xml version="1.0" encoding="UTF-8"?>
	<rss version="2.0"
		xmlns:content="http://purl.org/rss/1.0/modules/content/"
		xmlns:wfw="http://wellformedweb.org/CommentAPI/"
		xmlns:dc="http://purl.org/dc/elements/1.1/"
		xmlns:atom="http://www.w3.org/2005/Atom"
		xmlns:sy="http://purl.org/rss/1.0/modules/syndication/"
		xmlns:slash="http://purl.org/rss/1.0/modules/slash/"
		xmlns:georss="http://www.georss.org/georss"
		xmlns:geo="http://www.w3.org/2003/01/geo/wgs84_pos#"
		xmlns:media="http://search.yahoo.com/mrss/">
		<channel>
			<title>Republika Online - Nasional RSS Feed</title>
			<atom:link href="https://republika.co.id/rss/nasional" rel="self" type="application/rss+xml"/>
			<link>http://www.republika.co.id</link>
			<description></description>
			<lastBuildDate>Mon, 16 Mar 2020 14:29:53 +0700</lastBuildDate>
			<generator>http://www.republika.co.id/</generator>
			<language>id</language>
			<sy:updatePeriod>hourly</sy:updatePeriod>
			<sy:updateFrequency>1</sy:updateFrequency>
			<image>
				<url>https://static.republika.co.id/files/images/logo.png</url>
				<title>Republika Online - Nasional RSS Feed</title>
				<link>https://www.republika.co.id</link>
			</image>
			<item>
				<title>Dummy Title</title>
				<link>https://republika.co.id/berita/q79zht354/politikus-senior-yakin-amien-rais-tak-bentuk-pan-reformasi</link>
				<comments>https://republika.co.id/berita/q79zht354/politikus-senior-yakin-amien-rais-tak-bentuk-pan-reformasi</comments>
				<pubDate>Mon, 16 Mar 2020 14:29:53 +0700</pubDate>
				<dc:creator>Bayu Hermawan</dc:creator>
				<category>
					<![CDATA[Politik]]>
				</category>
				<media:content url="https://dummy.jpg" >
					<media:credit>Ist</media:credit>
					<media:title>Politikus PAN Tjatur Sapto Edy(Ist)</media:title>
				</media:content>
				<guid isPermaLink="false">https://republika.co.id/berita/q79zht354/politikus-senior-yakin-amien-rais-tak-bentuk-pan-reformasi</guid>
				<description><![CDATA[Dummy description]]></description>
				<content:encoded>
					<![CDATA[Dummy description]]>
				</content:encoded>
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
			Description: "Dummy description",
			PubDate:     1583988023,
		},
		news.News{
			Title:       "Sejarah bulu tangkis di Olimpiade: Mengapa Indonesia sulit lahirkan Susy Susanti generasi baru?",
			Url:         "http://www.bbc.com/indonesia/olahraga-51662063",
			Description: "Dummy description",
			PubDate:     1583980941,
		},
		news.News{
			Title:       "Virus corona: Karyawan apresiasi pembebasan pajak penghasilan, ekonom sebut 'perlu stimulus fiskal dan moneter' atasi perlambatan ekonomi",
			Url:         "http://www.bbc.com/indonesia/indonesia-51830029",
			Description: "Dummy description",
			PubDate:     1583976114,
		},
	}

	fetcher := NewFetcher()
	data, err := fetcher.Fetch(fmt.Sprintf("%s/bbc", server.URL))
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
			Title:        "Dummy Title",
			Url:          "https://www.antaranews.com/berita/1357722/surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota",
			MediaContent: "https://dummy.jpg",
			Description:  "Dummy description",
			PubDate:      1584262264,
		},
		news.News{
			Title:        "Dummy Title",
			Url:          "https://www.antaranews.com/foto/1357714/pencegahan-wabah-covid-19-di-kalimantan-tengah",
			MediaContent: "https://dummy.jpg",
			Description:  "Dummy description",
			PubDate:      1584262264,
		},
		news.News{
			Title:        "Dummy Title",
			Url:          "https://www.antaranews.com/video/1357690/presiden-imbau-masyarakat-bekerja-belajar-dan-beribadah-di-rumah",
			MediaContent: "https://dummy.jpg",
			Description:  "Dummy description",
			PubDate:      1584262264,
		},
	}

	fetcher := NewFetcher()
	url := fmt.Sprintf("%s/antaranews", server.URL)
	data, err := fetcher.Fetch(url)
	if err != nil {
		t.Fatalf("Error occured: %q", err)
	}

	assertLength(t, expectedResult, data)
	assertElements(t, expectedResult, data)
}

func TestDetikNewsFetcher(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fakeDetikServer))
	defer server.Close()

	expectedResult := []news.News{
		news.News{
			Title:        "Dummy Title",
			Url:          "https://news.detik.com/read/2020/03/15/182444/4940126/10/belajar-mengajar-tk-sampai-smp-di-kendari-pindah-ke-rumah-imbas-corona",
			MediaContent: "https://dummy.jpeg",
			Description:  "\"Dummy description\"",
			PubDate:      1584272086,
		},
	}

	fetcher := NewFetcher()
	url := fmt.Sprintf("%s/detik", server.URL)
	data, err := fetcher.Fetch(url)
	if err != nil {
		t.Fatalf("Error occured: %q", err)
	}

	assertLength(t, expectedResult, data)
	assertElements(t, expectedResult, data)
}

func TestOkezoneNewsFetcher(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fakeOkezoneServer))
	defer server.Close()

	expectedResult := []news.News{
		news.News{
			Title:        "Dummy Title",
			Url:          "https://megapolitan.okezone.com/read/2020/03/16/338/2184032/cegah-penyebaran-covid-19-pemkab-bekasi-tiadakan-kegiatan-publik",
			MediaContent: "https://dummy.jpg?w=300",
			Description:  "Dummy description",
			PubDate:      1584341598,
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

func TestRepublikaNewsFetcher(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fakeRepublikaServer))
	defer server.Close()

	expectedResult := []news.News{
		news.News{
			Title:        "Dummy Title",
			Url:          "https://republika.co.id/berita/q79zht354/politikus-senior-yakin-amien-rais-tak-bentuk-pan-reformasi",
			MediaContent: "https://dummy.jpg",
			Description:  "Dummy description",
			PubDate:      1584343793,
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
	t.Helper()
	for index, val := range expected {
		if val.Title != got[index].Title {
			t.Errorf("Struct value doesn't match!\nExpected: %v\nGot: %v\n", val.Title, got[index].Title)
		}

		if val.MediaContent != got[index].MediaContent {
			t.Errorf("Struct value doesn't match!\nExpected: %v\nGot: %v\n", val.MediaContent, got[index].MediaContent)
		}

		if val.Description != got[index].Description {
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
