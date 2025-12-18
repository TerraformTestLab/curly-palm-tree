package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	actual := HelloWorld()

	if actual != expected {
		t.Errorf("HelloWorld() = %q, want %q", actual, expected)
	}
}

func TestHelloWorldNotEmpty(t *testing.T) {
	result := HelloWorld()

	if result == "" {
		t.Error("HelloWorld() should not return an empty string")
	}
}
