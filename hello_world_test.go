package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}
}
