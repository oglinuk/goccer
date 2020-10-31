package goccer

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

var (
	pdfDir = "data/pdfs"
)

// HTTPCrawler for HTTP URLs
type HTTPCrawler struct {
	seed string
}

// NewHTTPCrawler constructor
func NewHTTPCrawler(s string) HTTPCrawler {
	return HTTPCrawler{
		seed: s,
	}
}

// Crawl c.seed and extract all URLs
func (c HTTPCrawler) Crawl() ([]string, error) {
	var collected []string

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 7,
	}

	resp, err := client.Get(c.seed)
	if err != nil {
		err = fmt.Errorf("crawlers::http.go::Crawl::client.Get(%s)::ERROR: %s", c.seed, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp == nil {
		err = errors.New("crawlers::http.go::Crawl::resp::NIL")
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		for _, URL := range c.extract(resp) {
			collected = append(collected, URL)
			//log.Printf("[COLLECTED] %s", URL)
		}
	} else {
		err = fmt.Errorf("crawlers::http.go::Crawl::resp.StatusCode(%d): %s", resp.StatusCode, c.seed)
		return nil, err
	}

	return collected, nil
}

func (c HTTPCrawler) extract(resp *http.Response) []string {
	if resp == nil {
		return nil
	}
	links := collectLinks(resp.Body)
	rebuiltLinks := []string{}

	for _, link := range links {
		url := rebuildURL(link, c.seed)
		if url != "" {
			rebuiltLinks = append(rebuiltLinks, url)
		}
	}

	return rebuiltLinks
}

func collectLinks(httpBody io.Reader) []string {
	links := make(map[string]struct{})
	col := []string{}
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			l := []string{}
			for k := range links {
				l = append(l, k)
			}
			return l
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(links, col)
				}
			}
		}
	}
}

func trimHash(l string) string {
	if strings.Contains(l, "#") {
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				return l[:n]
			}
		}
	}
	return l
}

func resolv(lm map[string]struct{}, ml []string) {
	for _, str := range ml {
		if _, exists := lm[str]; !exists {
			lm[str] = struct{}{}
		}
	}
}

func rebuildURL(href, base string) string {
	url, err := url.Parse(href)
	if err != nil {
		return ""
	}

	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}

	return baseURL.ResolveReference(url).String()
}
