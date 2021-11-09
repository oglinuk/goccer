package goccer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	actualWp = NewWorkerpool()
)


// TestNewWorkerpool checks that actualWp and all parts are not nil
func TestNewWorkerpool(t *testing.T) {
		assert := assert.New(t)
		assert.NotNil(t, actualWp)
		assert.NotNil(t, actualWp.jobs)
		assert.NotNil(t, actualWp.wg)
		assert.NotNil(t, actualWp.w)
		assert.NotNil(t, actualWp.c)
		assert.NotNil(t, actualWp.mu)
}

// TestQueue checks that the result of 3 URLs is not nil
func TestQueue(t *testing.T) {
	testURLs = []string{
		"https://fourohfournotfound.com",
		"https://en.wikipedia.org/wiki/Chaos_theory",
		"https://en.wikipedia.org/wiki/Deep_learning",
	}

	actualURLs := actualWp.Queue(testURLs)
	assert.NotNil(t, actualURLs)
}

