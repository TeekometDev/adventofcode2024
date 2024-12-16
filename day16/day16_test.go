package main

import (
	"teekometDev/filereader4"
	"testing"
)

func TestTask1_1(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test1.txt")
	got := Task1(file)
	want := 7036

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestTask1_2(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test2.txt")
	got := Task1(file)
	want := 11048

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}
