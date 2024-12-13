package main

import (
	"teekometDev/filereader5"
	"testing"
)

func TestT1(t *testing.T) {
	file := filereader5.ReadFile("test.txt")
	got := Task1(file)
	want := 480

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT2(t *testing.T) {
	file := filereader5.ReadFile("test.txt")
	got := Task2(file)
	want := -1

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}
