package main

import (
	"testing"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := Book{
		Title:  "The Deeper Love of Go",
		Author: "John Arundel",
		Copies: 15,
	}

	want := "The Deeper Love of Go by John Arundel - 15 copies"
	got := BookToString(input)
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}
