package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("aaron", func(t *testing.T) {
		res := Hello("Aaron")
		target := "Hello, Aaron"
		AssertEqual(t, res, target)
	})
	t.Run("empty string", func(t *testing.T) {
		res := Hello("")
		target := "Hello, World"
		AssertEqual(t, res, target)
	})
}

func AssertEqual(t testing.TB, result any, expected any) {
	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}
}
