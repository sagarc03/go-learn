package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sagar")
	want := "Hello, Sagar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
