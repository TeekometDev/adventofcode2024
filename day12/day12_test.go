package main

import (
	"teekometDev/filereader4"
	"testing"
)

func TestT1Ex1(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test1.txt")
	got := Task1(file)
	want := 140

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT1Ex2(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test2.txt")
	got := Task1(file)
	want := 772

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT1Ex3(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test3.txt")
	got := Task1(file)
	want := 1930

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT2Ex1(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test1.txt")
	got := Task2(file)
	want := 80

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT2Ex2(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test2.txt")
	got := Task2(file)
	want := 436

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT2Ex3(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test3.txt")
	got := Task2(file)
	want := 1206

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT2Ex4(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test4.txt")
	got := Task2(file)
	want := 368

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT2ExE(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test-e.txt")
	got := Task2(file)
	want := 236

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}
