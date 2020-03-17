package crawling

import (
	"github.com/imam98/api-wartapedia/pkg/news"
	"testing"
)

func TestGenSource(t *testing.T) {
	testcases := []struct{
		name string
		given news.SourceFlag
		expected string
	} {
		{
			name: "Source: AntaraNews",
			given: news.CAT_NASIONAL | news.ANTARANEWS,
			expected: "AntaraNews",
		},
		{
			name: "Source: BBC",
			given: news.CAT_NASIONAL | news.BBC,
			expected: "BBC",
		},
		{
			name: "Source: Detik",
			given: news.CAT_DUNIA | news.DETIK,
			expected: "Detik",
		},
		{
			name: "Source: Okezone",
			given: news.CAT_TEKNO | news.OKEZONE,
			expected: "Okezone",
		},
		{
			name: "Source: Republika",
			given: news.CAT_TEKNO | news.REPUBLIKA,
			expected: "Republika",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := parseSourceFromFlags(tc.given)
			if got != tc.expected {
				t.Errorf("Value doesn't match\nExpected: %q\nGiven: %b\nGot: %q\n", tc.expected, tc.given, got)
			}
		})
	}
}

func TestGenDocID(t *testing.T) {
	testcases := []struct{
		name string
		given string
		flags news.SourceFlag
		expected string
	} {
		{
			name: "Source: AntaraNews",
			given: "https://www.antaranews.com/berita/1357722/surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota",
			flags: news.CAT_NASIONAL | news.ANTARANEWS,
			expected: "atn::surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota",
		},
		{
			name: "Source: BBC",
			given: "http://www.bbc.com/indonesia/olahraga-51662063",
			flags: news.CAT_DUNIA | news.BBC,
			expected: "bbc::olahraga-51662063",
		},
		{
			name: "Source: Okezone",
			given: "https://megapolitan.okezone.com/read/2020/03/16/338/2184032/cegah-penyebaran-covid-19-pemkab-bekasi-tiadakan-kegiatan-publik",
			flags: news.CAT_DUNIA | news.OKEZONE,
			expected: "okz::cegah-penyebaran-covid-19-pemkab-bekasi-tiadakan-kegiatan-publik",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, prefix := parseSourceFromFlags(tc.flags)
			got := genDocID(prefix, tc.given)

			if got != tc.expected {
				t.Errorf("Value doesn't match\nExpected: %q\nGiven: %q\nWith Flags: %b\nGot: %q\n", tc.expected, tc.given, tc.flags, got)
			}
		})
	}
}