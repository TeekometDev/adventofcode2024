package main

import "testing"

func TestT1(t *testing.T) {
	got := Task1("test.txt")
	want := 18

	if got != want {
		t.Errorf("got %d instead of %d", got, want)
	}
}

func TestT1Input(t *testing.T) {
	got := Task1("input.txt")
	want := 2493

	if got != want {
		t.Errorf("got %d instead of %d", got, want)
	}
}

func TestT2(t *testing.T) {
	got := Task2("test.txt")
	want := 9

	if got != want {
		t.Errorf("got %d instead of %d", got, want)
	}
}
