package main

import (
	"teekometDev/filereader4"
	"testing"
)

func TestT1(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test.txt")
	got := Task1(file)
	want := 36

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func TestT2(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test.txt")
	got := Task2(file)
	want := 81

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}
