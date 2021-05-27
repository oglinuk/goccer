package goccer

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert := assert.New(t)
		assert.NotNil(t, actualWp)
		assert.NotNil(t, actualWp.jobs)
		assert.NotNil(t, actualWp.wg)
		assert.NotNil(t, actualWp.w)
}

func TestQueue(t *testing.T) {
	actualURLs := actualWp.Queue(testURLs)
	assert.NotNil(t, actualURLs)
}

