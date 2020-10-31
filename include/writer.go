package goccer

import (
	"fmt"
	"net/url"
)

// MemoryPool is a domain and a map containing routes of the domain
type MemoryPool struct {
	roots map[string]map[string]struct{}
}

// NewMemoryPool constructor
func NewMemoryPool() *MemoryPool {
	return &MemoryPool{
		roots: make(map[string]map[string]struct{}),
	}
}

// Write path domains/routes if not existing already
func (mp *MemoryPool) Write(paths []string) error {
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

func (mp *MemoryPool) GetRoots() []string {
	var paths []string
	for domain := range mp.roots {
		for route := range mp.roots[domain] {
			path := fmt.Sprintf("%s%s", domain, route)
			paths = append(paths, path)
		}
	}

	return paths
}
