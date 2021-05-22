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
	Client *http.Client
	Err error
	Root string
}

// NewCrawler constructor
func NewCrawler(r string) *crawler {
	return &crawler{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: time.Second * 7,
		},
		Err: nil,
		Root: r,
	}
}

// Crawl the crawlers Root and returns the extracted URLs
func (c *crawler) Crawl() []string {
	resp, err := c.Client.Get(c.Root)
	if err != nil {
		c.Err = err
		return nil
	}
	defer resp.Body.Close()

	return c.ParseHTML(resp.Body)
}

// ParseHTML takes an io.Reader (http.Response.Body), extracts all
// <a>nchor tags, and returns all rebuilt <a> tags as full URLs
func (c *crawler) ParseHTML(body io.Reader) []string {
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
					rebuilt := c.RebuildURL(attr.Val)
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

// RebuildURL returns a string depending on the state of href
func (c *crawler) RebuildURL(href string) string {
	var rebuilt string

	// check if href is already a valid URL
	if strings.HasPrefix(href, "http") {
		rebuilt = href
	}

	// check if href has '//' prefix
	if href[0] == '/' && href[1] == '/' {
		rebuilt = fmt.Sprintf("https:%s", href)
	}

	// check if href has '/' prefix
	if href[0] == '/' && href[1] != '/' {
		href = strings.TrimPrefix(href, "/")
		rebuilt = fmt.Sprintf("%s/%s", c.Root, href)
	}

	// check if href has '#' prefix
	if href[0] == '#' {
		rebuilt = fmt.Sprintf("%s%s", c.Root, href)
	}

	return rebuilt
}
