package elasticsearch

import (
	"bytes"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/imam98/api-wartapedia/pkg/news"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

var (
	fixtures = make(map[string]io.ReadCloser)
)

func init() {
	fixtureFiles, err := filepath.Glob("testdata/*.json")
	if err != nil {
		panic(fmt.Sprintf("Cannot glob fixture files: %s", err))
	}

	for _, fpath := range fixtureFiles {
		f, err := ioutil.ReadFile(fpath)
		if err != nil {
			panic(fmt.Sprintf("Cannot read fixture file: %s", err))
		}
		fixtures[filepath.Base(fpath)] = ioutil.NopCloser(bytes.NewReader(f))
	}
}

func fixture(fname string) io.ReadCloser {
	out := new(bytes.Buffer)
	b1 := bytes.NewBuffer([]byte{})
	b2 := bytes.NewBuffer([]byte{})
	tr := io.TeeReader(fixtures[fname], b1)

	defer func() { fixtures[fname] = ioutil.NopCloser(b1) }()
	io.Copy(b2, tr)
	out.ReadFrom(b2)

	return ioutil.NopCloser(out)
}

type fakeTransport struct {
	Response    *http.Response
	RoundTripFn func(req *http.Request) (*http.Response, error)
}

func (ft *fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return ft.RoundTripFn(req)
}

func TestNewRepository(t *testing.T) {
	given := Config{
		Client:    nil,
		IndexName: "",
	}
	expected := &repository{
		client:    nil,
		indexName: "",
	}

	got := NewRepository(given)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Unexpected value\nExpected: %#v\nGot: %#v\n", expected, got)
	}
}

func TestFind(t *testing.T) {
	ft := fakeTransport{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixture("find_result.json"),
		},
	}
	ft.RoundTripFn = func(req *http.Request) (response *http.Response, err error) {
		path := strings.Split(req.URL.Path, "/")
		if path[len(path)-1] != "abc:123" {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       fixture("find_not_found.json"),
			}, nil
		}

		return ft.Response, nil
	}

	client, err := es.NewClient(es.Config{
		Transport: &ft,
	})
	if err != nil {
		t.Fatalf("Expect no error, got: %q", err.Error())
	}

	repo := NewRepository(Config{
		Client:    client,
		IndexName: "testing",
	})

	t.Run("Document exists", func(t *testing.T) {
		expected := news.News{
			ID:           "abc:123",
			Source:       "abc",
			Title:        "Dummy Title",
			MediaContent: "http://dummy.jpg",
			Url:          "http://dummy.id",
			Description:  "Dummy description",
			PubDate:      1585901013,
		}

		got, err := repo.Find("abc:123")
		if err != nil {
			t.Errorf("Expect no error, got: %q", err.Error())
		}

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Unexpected value\nExpected: %#v\nGot: %#v", expected, got)
		}
	})

	t.Run("Document not exists", func(t *testing.T) {
		expected := news.ErrItemNotFound
		_, got := repo.Find("abc:235")
		if got == nil {
			t.Error("Expect to get error, got no error")
		} else if expected != got {
			t.Errorf("Unexpected error\nExpected: %q\nGot: %q\n", expected, got)
		}
	})
}
