package goccer

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	actualMp = NewMemorypool()
)

func TestNewMemorypool(t *testing.T) {
	assert.NotNil(t, actualMp)
	assert.NotNil(t, actualMp.mapping)
	assert.NotNil(t, actualMp.mapping["error"])
}

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
