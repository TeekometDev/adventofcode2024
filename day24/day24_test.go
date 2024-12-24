package main

import (
	"teekometDev/filereader5"
	"testing"
)

func TestTask1Small(t *testing.T) {
	file := filereader5.ReadFile("test1.txt")
	got := Task1(file)
	want := 4

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestTask1Big(t *testing.T) {
	file := filereader5.ReadFile("test2.txt")
	got := Task1(file)
	want := 2024

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}
