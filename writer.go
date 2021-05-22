package goccer

import (
	"fmt"
	"net/url"
)

type memoryPool struct {
	// map[domain]map[route]struct{}
	// map[https://golang.org]map[/pkg/net/http]struct{}
	mapping map[string]map[string]struct{}
}

func NewMemorypool() *memoryPool {
	mp := &memoryPool{
		mapping: make(map[string]map[string]struct{}),
	}

	// Initialize a new map for the 'error' key value to allow us to store
	// URLs that errored
	mp.mapping["error"] = make(map[string]struct{})

	return mp
}

func (mp *memoryPool) write(paths []string) error {
	for _, p := range paths {
		parsed, err := url.Parse(p)
		if err != nil {
			return err
		}

		// Extract scheme (http, https, ...)
		scheme := parsed.Scheme

		// Extract domain (golang.org, fourohfournotfound.com, ...)
		domain := parsed.Hostname()

		// If the domain fails, track the URL under "error" key
		if domain == "" || domain == " " {
			domain = "error"
			mp.mapping[domain][parsed.String()] = struct{}{}
		}

		domain = fmt.Sprintf("%s://%s", scheme, domain)
		if _, exists := mp.mapping[domain]; !exists {
			mp.mapping[domain] = make(map[string]struct{})
		}

		if len(parsed.Path) > 0 {
			if _, exists := mp.mapping[domain][parsed.Path]; !exists {
				mp.mapping[domain][parsed.Path] = struct{}{}
			}
		}
	}

	return nil
}

// GetPaths returns all paths stored in mp.mapping after rebuilding them
func (mp *memoryPool) GetPaths() []string {
	var paths []string

	// TODO: Need to do something else with mp.mapping["errors"]
	for k := range mp.mapping {
		if k != "error" {
			for v := range mp.mapping[k] {
				paths = append(paths, fmt.Sprintf("%s%s", k, v))
			}
		}
	}

	return paths
}
