package goccer

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"io"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// TODO: Abstract below to allow for different types of crawlers (ie: db, fs)
type crawler struct {
	client *http.Client
	seed string
}

// NewCrawler constructor
func NewCrawler() *crawler {
	return &crawler{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: time.Second * 15,
		},
	}
}

// Crawl the given seed and return the URLs from c.parseHTML
func (c *crawler) Crawl(seed string) ([]string, error) {
	if seed == "" || seed == " " {
		return nil, nil
	}	else {
		c.seed = seed
	}

	resp, err := c.client.Get(c.seed)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return c.parseHTML(resp.Body), nil
}

// parseHTML takes an io.Reader (http.Response.Body), extracts all
// <a>nchor tags, and returns all rebuilt <a> tags as full URLs
func (c *crawler) parseHTML(body io.Reader) []string {
	if body == nil {
		return nil
	}

	var parsed []string
	checked := make(map[string]struct{})

	// https://pkg.go.dev/golang.org/x/net/html
	page := html.NewTokenizer(body)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return parsed
		}

		token := page.Token()
		// Example token.DataAtom possibilities: h1, p, code, a, ul, li, ...
		// We only care about 'a' though
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				// attr can be a valid URL or a route
				// ie - "https://github.com/afkworks/spec-kn" || "books" || "#cite_note-77"
				if attr.Key == "href" && attr.Val != "" {
					rebuilt := c.rebuildURL(attr.Val)
					if _, exists := checked[rebuilt]; !exists {
						parsed = append(parsed, rebuilt)
						checked[rebuilt] = struct{}{}
					}
				}
			}
		}
	}

	return parsed
}

// rebuildURL returns a string depending on the state of href
func (c *crawler) rebuildURL(href string) string {
	rebuilt := ""

	// check if href is already a valid URL
	if strings.HasPrefix(href, "http") {
		rebuilt = href
	}

	// check if href iS '/', '//', or '#'
	if len(href) < 3 {
		rebuilt = c.seed
	}

	// if we get to this point, rebuild using c.seed and href
	if rebuilt == "" {
		if strings.HasPrefix(href, "//") {
			rebuilt = fmt.Sprintf("https:%s", href)
		} else if strings.HasPrefix(href, "/") || strings.HasPrefix(href, "#") {
			rebuilt = fmt.Sprintf("%s%s", c.seed, href)
		} else {
			rebuilt = fmt.Sprintf("%s/%s", c.seed, href)
		}
	}

	return rebuilt
}
