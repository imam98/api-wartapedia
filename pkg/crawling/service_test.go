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
			got := genSourceFromFlags(tc.given)
			if got != tc.expected {
				t.Errorf("Value doesn't match\nExpected: %q\nGiven: %b\nGot: %q\n", tc.expected, tc.given, got)
			}
		})
	}
}