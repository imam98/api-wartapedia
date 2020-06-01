package crawling

import (
	"github.com/imam98/api-wartapedia/pkg/domain"
	"testing"
)

func TestGenSource(t *testing.T) {
	testcases := []struct {
		name     string
		given    domain.RepoFlag
		expected string
	}{
		{
			name:     "Source: AntaraNews",
			given:    domain.CAT_NASIONAL | domain.ANTARANEWS,
			expected: "AntaraNews",
		},
		{
			name:     "Source: BBC",
			given:    domain.CAT_NASIONAL | domain.BBC,
			expected: "BBC",
		},
		{
			name:     "Source: Detik",
			given:    domain.CAT_DUNIA | domain.DETIK,
			expected: "Detik",
		},
		{
			name:     "Source: Okezone",
			given:    domain.CAT_TEKNO | domain.OKEZONE,
			expected: "Okezone",
		},
		{
			name:     "Source: Republika",
			given:    domain.CAT_TEKNO | domain.REPUBLIKA,
			expected: "Republika",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.given.SourceString()
			if got != tc.expected {
				t.Errorf("Value doesn't match\nExpected: %q\nGiven: %b\nGot: %q\n", tc.expected, tc.given, got)
			}
		})
	}
}

func TestGenDocID(t *testing.T) {
	testcases := []struct {
		name     string
		given    string
		flags    domain.RepoFlag
		expected string
	}{
		{
			name:     "Source: AntaraNews",
			given:    "https://www.antaranews.com/berita/1357722/surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota",
			flags:    domain.CAT_NASIONAL | domain.ANTARANEWS,
			expected: "atn::surabaya-belum-perlu-lockdown-antisipasi-covid-19-sebut-wali-kota",
		},
		{
			name:     "Source: BBC",
			given:    "http://www.bbc.com/indonesia/olahraga-51662063",
			flags:    domain.CAT_DUNIA | domain.BBC,
			expected: "bbc::olahraga-51662063",
		},
		{
			name:     "Source: Okezone",
			given:    "https://megapolitan.okezone.com/read/2020/03/16/338/2184032/cegah-penyebaran-covid-19-pemkab-bekasi-tiadakan-kegiatan-publik",
			flags:    domain.CAT_DUNIA | domain.OKEZONE,
			expected: "okz::cegah-penyebaran-covid-19-pemkab-bekasi-tiadakan-kegiatan-publik",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			prefix := parsePrefixFromFlags(tc.flags)
			got := genDocID(prefix, tc.given)

			if got != tc.expected {
				t.Errorf("Value doesn't match\nExpected: %q\nGiven: %q\nWith Flags: %b\nGot: %q\n", tc.expected, tc.given, tc.flags, got)
			}
		})
	}
}
