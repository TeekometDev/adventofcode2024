package main

import (
	"teekometDev/filereader5"
	"testing"
)

func TestTask1(t *testing.T) {
	file := filereader5.ReadFile("test.txt")
	got := Task1(file)
	want := 6

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestTask2(t *testing.T) {
	file := filereader5.ReadFile("test.txt")
	got := Task2(file)
	want := 16

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}
