package goccer

import (
	"crypto/tls"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

const (
	expectedRoot = "https://en.wikipedia.org/wiki/Chaos_theory"
	testHTML = "testdata/chaos-theory.html"
)

var (
	c = NewCrawler(expectedRoot)
)

func TestNewCrawler(t *testing.T) {
	if c == nil {
		t.Errorf("Expected: crawler; Got: nil")
	}

	if c.Err != nil {
		t.Errorf("Expected: nil; Got: %s", c.Err.Error())
	}

	if c.Root != expectedRoot {
		t.Errorf("Expected: %s; Got: %s", expectedRoot, c.Root)
	}

	expectedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 7,
	}
	if !reflect.DeepEqual(c.Client, expectedClient) {
		t.Errorf("Expected: %v; Got: %v", expectedClient, c.Client)
	}
}

func TestCrawl(t *testing.T) {
	_ = c.Crawl()
	if c.Err != nil {
		t.Errorf("Expected: nil; Got: %s", c.Err.Error())
	}
}

func TestParseHTML(t *testing.T) {
	testHTML, err := os.Open(testHTML)
	if err != nil {
		t.Errorf("Expected: nil; Got: %s", err.Error())
	}
	defer testHTML.Close()

	parsed := c.ParseHTML(testHTML)
	if c.Err != nil {
		t.Errorf("Expected: nil; Got: %s", c.Err.Error())
	}

	expectedLen := 1397
	if len(parsed) != expectedLen {
		t.Errorf("Expected: %d; Got: %d", expectedLen, len(parsed))
	}
}

func TestRebuildURL(t *testing.T) {
	hrefs := []string{
		"https://github.com/afkworks/spec-kn",
		"/wiki/File:Double-compound-pendulum.gif",
		"#cite_note-77",
		"/w/index.php?title=Chaos_theory&action=edit&section=6",
		"/wiki/File:Logistic_Map_Bifurcation_Diagram,_Matplotlib.svg",
		"//en.wikipedia.org/wiki/Chaos_theory",
	}


	expectedURLs := []string{
		"https://github.com/afkworks/spec-kn",
		"https://en.wikipedia.org/wiki/Chaos_theory/wiki/File:Double-compound-pendulum.gif",
		"https://en.wikipedia.org/wiki/Chaos_theory#cite_note-77",
		"https://en.wikipedia.org/wiki/Chaos_theory/w/index.php?title=Chaos_theory&action=edit&section=6",
		"https://en.wikipedia.org/wiki/Chaos_theory/wiki/File:Logistic_Map_Bifurcation_Diagram,_Matplotlib.svg",
		"https://en.wikipedia.org/wiki/Chaos_theory",
	}

	var actualURLs []string

	for _, href := range hrefs {
		actualURLs = append(actualURLs, c.RebuildURL(href))
	}

	for i, URL := range actualURLs {
		if URL != expectedURLs[i] {
			t.Errorf("\nExpected: %s\nGot: %s", expectedURLs[i], URL)
		}
	}
}
