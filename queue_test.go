package goccer

import (
	"testing"
)

var (
	actualWp = NewWorkerpool()

	testURLs = []string{
		"https://fourohfournotfound.com",
		"https://en.wikipedia.org/wiki/Chaos_theory",
		"https://en.wikipedia.org/wiki/Deep_learning",
	}
)

func TestNewWorkerpool(t *testing.T) {
	if actualWp == nil {
		t.Error("Expected: something; Got: nil")
	}
}

func TestQueue(t *testing.T) {
	if len(actualWp.Queue(testURLs)) < 1 {
		t.Error("Expected > 1; Got < 1")
	}
}

