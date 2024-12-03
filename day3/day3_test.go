package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q instead of %q", got, want)
	}
}

func TestT1(t *testing.T) {
	got := SolveT1("test.txt")
	want := 161

	if got != want {
		t.Errorf("got %d instead of %d", got, want)
	}
}

func TestT2(t *testing.T) {
	got := SolveT2("test2.txt")
	want := 48

	if got != want {
		t.Errorf("got %d instead of %d", got, want)
	}
}
