package main

import (
	"teekometDev/filereader5"
	"testing"
)

func TestT1(t *testing.T) {
	text := filereader5.ReadFile("test.txt")
	got := Task1(text, Playground{x_dim: 11, y_dim: 7})
	want := 12

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}
