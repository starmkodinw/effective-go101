package test

import (
	"testing"
)

func TestFGithub(t *testing.T) {
	v := "Hello"
	expected := "World" // Replace with the expected result after calling fGithub

	result := fGithub(v)

	if result != expected {
		t.Errorf("fGithub(v) failed, expected %v, got %v", expected, result)
	}
}
