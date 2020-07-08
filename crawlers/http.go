package crawlers

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/oglinuk/goccer"

	"golang.org/x/net/html"
)

// HTTPCrawler for HTTP URLs
type HTTPCrawler struct {
	seed   string
	writer goccer.Writer
}

// NewHTTPCrawler constructor
func NewHTTPCrawler(s string, w goccer.Writer) HTTPCrawler {
	return HTTPCrawler{
		seed:   s,
		writer: w,
	}
}

// Crawl c.seed and extract all URLs
func (c HTTPCrawler) Crawl() {
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
		log.Printf("crawlers::http.go::client.Get(%s)::ERROR: %s", c.seed, err.Error())
	}

	if resp == nil {
		log.Println("crawlers::http.go::resp::NIL")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		for _, URL := range c.extract(resp) {
			// TODO: Replace below with refactored http writer implementation
			log.Println(URL)
		}
	} else {
		log.Printf("crawlers::http.go::resp.StatusCode: %d", resp.StatusCode)
	}
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

	resp.Body.Close()

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
