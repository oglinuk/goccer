package goccer

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	actualMp = NewMemorypool()
)

// TestNewMemorypool ensures that actualMp and its parts are not nil
func TestNewMemorypool(t *testing.T) {
	assert.NotNil(t, actualMp)
	assert.NotNil(t, actualMp.mapping)
	assert.NotNil(t, actualMp.mapping["error"])
}

// TestWrite ensures that actualWrittenURLs and expectedMap are equal
func TestWrite(t *testing.T) {
	actualWrittenURLs := actualMp.write(testURLs)
	assert.Nil(t, actualWrittenURLs)

	expectedMap := map[string]map[string]struct{}{
		"https://fourohfournotfound.com": make(map[string]struct{}),
		"https://en.wikipedia.org": map[string]struct{}{
			"/wiki/Chaos_theory": struct{}{},
			"/wiki/Deep_learning": struct{}{},
		},
		"error": make(map[string]struct{}),
	}
	assert.Equal(t, expectedMap, actualMp.mapping)
}

// TestGetPaths ensures that actualPaths is equal to expectedPaths
func TestGetPaths(t *testing.T) {
	actualMp.write(testURLs)

	actualPaths := actualMp.GetPaths()
	assert.NotNil(t, actualPaths)

	expectedPaths := []string{
		"https://fourohfournotfound.com",
		"https://en.wikipedia.org",
		"https://en.wikipedia.org/wiki/Chaos_theory",
		"https://en.wikipedia.org/wiki/Deep_learning",
	}

	sort.Strings(expectedPaths)
	sort.Strings(actualPaths)

	assert.Equal(t, expectedPaths, actualPaths)
}
