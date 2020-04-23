package listing

import (
	"github.com/imam98/api-wartapedia/pkg/news"
	"reflect"
	"testing"
)

type fakeFetcher struct{}

func (f *fakeFetcher) Fetch(url string) ([]news.News, error) {
	data := []news.News{
		news.News{
			Title:        "Dummy Title",
			MediaContent: "http://dummy.jpg",
			Url:          "http://dummy.id",
			Description:  "Dummy description",
			PubDate:      1583049181,
		},
	}

	return data, nil
}

func TestNewService(t *testing.T) {
	fetcher := &fakeFetcher{}
	expected := &listing{nf: fetcher}
	got := NewService(fetcher)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Unexpected service\nExpected: %#v\nGot: %#v\n", expected, got)
	}
}

func TestGetNews(t *testing.T) {
	t.Run("Valid source", func(t *testing.T) {
		fetcher := &fakeFetcher{}
		service := NewService(fetcher)

		expected := []news.News{
			news.News{
				Title:        "Dummy Title",
				MediaContent: "http://dummy.jpg",
				Url:          "http://dummy.id",
				Description:  "Dummy description",
				PubDate:      1583049181,
			},
		}
		got, err := service.GetNews(news.CAT_NASIONAL | news.DETIK)
		assertError(t, nil, err)

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Unexpected value\nExpected: %#v\nGot: %#v\n", expected, got)
		}
	})

	t.Run("Invalid source", func(t *testing.T) {
		fetcher := &fakeFetcher{}
		service := NewService(fetcher)

		_, got := service.GetNews(news.CAT_TEKNO | news.DETIK)
		assertError(t, news.ErrSourceNotFound, got)
	})
}

func TestGetSources(t *testing.T) {
	testcases := []struct {
		name     string
		given    news.RepoFlag
		expected []string
	}{
		{
			name:     "Category: Nasional",
			given:    news.CAT_NASIONAL,
			expected: []string{"AntaraNews", "BBC", "Detik", "Okezone", "Republika"},
		},
		{
			name:     "Category: Dunia",
			given:    news.CAT_DUNIA,
			expected: []string{"AntaraNews", "BBC", "Detik", "Republika"},
		},
		{
			name:     "Category: Tekno",
			given:    news.CAT_TEKNO,
			expected: []string{"AntaraNews", "Okezone", "Republika"},
		},
	}

	fetcher := &fakeFetcher{}
	service := NewService(fetcher)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := service.GetSourcesFromCategory(tc.given)
			assertError(t, nil, err)

			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("Unexpected value\nExpected: %v\nGiven: %b\nGot: %v\n", tc.expected, tc.given, got)
			}
		})
	}

	t.Run("Category: Invalid", func(t *testing.T) {
		given := news.RepoFlag(news.ANTARANEWS)
		expected := ErrInvalidCategoryFlag

		_, got := service.GetSourcesFromCategory(given)
		assertError(t, expected, got)
	})

	t.Run("Category: Invalid #2", func(t *testing.T) {
		given := news.RepoFlag(news.CAT_NASIONAL | news.ANTARANEWS)
		expected := ErrInvalidCategoryFlag

		_, got := service.GetSourcesFromCategory(given)
		assertError(t, expected, got)
	})
}

func BenchmarkGetSources(b *testing.B) {
	fetcher := &fakeFetcher{}
	service := NewService(fetcher)
	for i := 0; i < b.N; i++ {
		service.GetSourcesFromCategory(news.RepoFlag(news.CAT_TEKNO))
	}
}

func assertError(t *testing.T, expected error, got error) {
	t.Helper()

	if got == nil && expected != nil {
		t.Error("Expect to get error")
	} else if got != nil && expected == nil {
		t.Fatalf("Expect no error, got: %q\n", got)
	} else if got != nil && expected != nil && expected != got {
		t.Errorf("Unexpected error\nExpected: %q\nGot: %q", expected, got)
	}
}
