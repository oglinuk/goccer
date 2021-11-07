package goccer

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	expectedSeed = "https://en.wikipedia.org/wiki/Chaos_theory"
	testHTML = "testdata/chaos-theory.html"
)

var (
	c = NewCrawler()
)

func TestNewCrawler(t *testing.T) {
	/*
	if c == nil {
		t.Errorf("Expected: crawler; Got: nil")
	}
	*/
	assert.NotNil(t, c)

	expectedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 7,
	}

	/*
	if !reflect.DeepEqual(c.Client, expectedClient) {
		t.Errorf("Expected: %v; Got: %v", expectedClient, c.Client)
	}
	*/
	assert.Equal(t, expectedClient, c.Client)
}

func TestCrawl(t *testing.T) {
	_, err := c.Crawl("")
	assert.Nil(t, err)

	_, err = c.Crawl(expectedSeed)
	assert.Nil(t, err)
}

func TestParseHTML(t *testing.T) {
	assert.Nil(t, c.ParseHTML(nil))

	testHTML, err := os.Open(testHTML)
	/*
	if err != nil {
		t.Errorf("Expected: nil; Got: %s", err.Error())
	}
	*/
	assert.Nil(t, err)
	defer testHTML.Close()

	parsed := c.ParseHTML(testHTML)

	/*
	expectedLen := 1397
	if len(parsed) != expectedLen {
		t.Errorf("Expected: %d; Got: %d", expectedLen, len(parsed))
	}
	*/
	assert.Equal(t, 1397, len(parsed))
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
		/*
		if URL != expectedURLs[i] {
			t.Errorf("\nExpected: %s\nGot: %s", expectedURLs[i], URL)
		}
		*/
		assert.Equal(t, expectedURLs[i], URL)
	}
}
