package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	in := "dood"
	want := "poop"

	got := reverseString(in)

	if got != want {
		t.Errorf("reverseString(%q) == %q, want %q", in, got, want)
	}
}

func TestFlipText(t *testing.T) {
	in := "a"
	want := "ɐ"

	got := flipText(in)

	if got != want {
		t.Errorf("flipText(%q) == %q, want %q", in, got, want)
	}
}
