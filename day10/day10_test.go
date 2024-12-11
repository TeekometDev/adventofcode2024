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

func TestPart1Input(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("input.txt")
	got := Task1(file)
	want := 794

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	file := filereader4.ReadFileAsMatrix("input.txt")
	for n := 0; n < b.N; n++ {
		Task1(file)
	}
}

func BenchmarkPart1Challenge(b *testing.B) {
	file := filereader4.ReadFileAsMatrix("challenge.txt")
	for n := 0; n < b.N; n++ {
		Task1(file)
	}
}

func TestPart2Input(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("input.txt")
	got := Task2(file)
	want := 1706

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func BenchmarkPart2(b *testing.B) {
	file := filereader4.ReadFileAsMatrix("input.txt")
	for n := 0; n < b.N; n++ {
		Task2(file)
	}
}

func BenchmarkPart2Challenge(b *testing.B) {
	file := filereader4.ReadFileAsMatrix("challenge.txt")
	for n := 0; n < b.N; n++ {
		Task2(file)
	}
}
