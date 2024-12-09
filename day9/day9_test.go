package main

import (
	"teekometDev/filereader3"
	"testing"
)

func TestT1(t *testing.T) {
	line := filereader3.ReadFile("test.txt")
	got := Task1(line)
	want := 1928

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func TestT2(t *testing.T) {
	line := filereader3.ReadFile("test.txt")
	got := Task2(line)
	want := 2858

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}
