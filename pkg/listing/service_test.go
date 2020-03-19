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
			MediaContent: news.Media{Src: "http://dummy.jpg"},
			Url:          "http://dummy.id",
			Description:  news.Description{Text: "Dummy description"},
			PubDate:      "01 Mar 2020 14:53:01 +0700",
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
				MediaContent: news.Media{Src: "http://dummy.jpg"},
				Url:          "http://dummy.id",
				Description:  news.Description{Text: "Dummy description"},
				PubDate:      "01 Mar 2020 14:53:01 +0700",
			},
		}
		got, err := service.GetNews(news.CAT_NASIONAL | news.DETIK)
		if err != nil {
			t.Fatalf("Expect no error, got: %q\n", err)
		}

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Unexpected value\nExpected: %#v\nGot: %#v\n", expected, got)
		}
	})

	t.Run("Invalid source", func(t *testing.T) {
		fetcher := &fakeFetcher{}
		service := NewService(fetcher)

		_, err := service.GetNews(news.CAT_TEKNO | news.DETIK)
		if err == nil {
			t.Fatal("Should throw error here")
		}

		if err != ErrSourceNotFound {
			t.Errorf("Unexpected error\nExpected: %q\nGot: %q\n", ErrSourceNotFound, err)
		}
	})
}
