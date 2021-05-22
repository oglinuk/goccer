package goccer

import (
	"testing"
)

var (
	actualMp := NewMemorypool()
)

func TestNewMemorypool(t *testing.T) {
	if actualMp == nil {
		t.Error("Expected: something; Got: nil")
	}
}

func TestWrite(t *testing.T) {
	if err := actualMp.write(testURLs); err != nil {
		t.Errorf("Expected: nil; Got: %s", err.Error())
	}
}

func TestGetPaths(t *testing.T) {
	actualMp.write(testURLs)

	if len(actualMp.GetPaths()) < 1 {
		t.Error("Expected: something; Got: nil")
	}
}
