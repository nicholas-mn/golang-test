package main

import (
	"bytes"
	"testing"
)

func TestGrees(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nicholas")

	got := buffer.String()
	want := "Hello, Nicholas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
