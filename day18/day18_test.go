package main

import (
	"teekometDev/filereader5"
	"testing"
)

func TestTask1(t *testing.T) {
	file := filereader5.ReadFile("test.txt")
	got := Task1(file, 6, 6, 12)
	want := 22

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestTask2(t *testing.T) {
	file := filereader5.ReadFile("test.txt")
	gotx, goty := Task2(file, 6, 6)
	wantx := 6
	wanty := 1

	if gotx != wantx || goty != wanty {
		t.Errorf("Got %d, %d and wanted %d, %d\n", gotx, goty, wantx, wanty)
	}
}
