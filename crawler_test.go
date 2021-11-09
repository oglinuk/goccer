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

// TestNewCrawler checks to make sure the crawler is not nil and if it is
// equal to the expectedClient
func TestNewCrawler(t *testing.T) {
	assert.NotNil(t, c)

	expectedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 15,
	}

	assert.Equal(t, expectedClient, c.client)
}

// TestCrawl checks that err is nil for both "" and expectedSeed
func TestCrawl(t *testing.T) {
	_, err := c.Crawl("")
	assert.Nil(t, err)

	_, err = c.Crawl(expectedSeed)
	assert.Nil(t, err)
}

// TestParseHTML checks that passing nil to c.parseHTML returns nil and
// that the testHTML file returns 1397 links
func TestParseHTML(t *testing.T) {
	assert.Nil(t, c.parseHTML(nil))

	testHTML, err := os.Open(testHTML)
	assert.Nil(t, err)
	defer testHTML.Close()

	parsed := c.parseHTML(testHTML)
	assert.Equal(t, 1397, len(parsed))
}

// TestRebuildURL checks that all hrefs are returned as its expected URL
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
		actualURLs = append(actualURLs, c.rebuildURL(href))
	}

	for i, URL := range actualURLs {
		assert.Equal(t, expectedURLs[i], URL)
	}
}
