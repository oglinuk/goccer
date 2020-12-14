package goccer

import (
	"fmt"
	"net/url"
)

// memoryPool is a domain and a map containing routes of the domain
type memoryPool struct {
	roots map[string]map[string]struct{}
}

// newMemoryPool constructor
func newMemoryPool() *memoryPool {
	return &memoryPool{
		roots: make(map[string]map[string]struct{}),
	}
}

// write path domains/routes if not existing already
func (mp *memoryPool) write(paths []string) error {
	for _, p := range paths {
		u, err := url.Parse(p)
		if err != nil {
			return err
		}

		domain := u.Hostname()
		if domain == "" || domain == " " {
			domain = "error"
		}

		if _, exists := mp.roots[domain]; !exists {
			mp.roots[domain] = make(map[string]struct{})
		}
		if _, exists := mp.roots[domain][u.Path]; !exists {
			mp.roots[domain][u.Path] = struct{}{}
		}
	}

	return nil
}

func (mp *memoryPool) getRoots() []string {
	var paths []string
	for domain := range mp.roots {
		for route := range mp.roots[domain] {
			path := fmt.Sprintf("%s%s", domain, route)
			paths = append(paths, path)
		}
	}

	return paths
}
