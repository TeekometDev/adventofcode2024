package main

import (
	"teekometDev/filereader4"
	"testing"
)

func TestT1(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test.txt")
	got := Task1(file)
	want := 14

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func TestT2(t *testing.T) {
	file := filereader4.ReadFileAsMatrix("test.txt")
	got := Task2(file)
	want := 34

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func TestCalcAntinode(t *testing.T) {
	obj1 := Coordinates{row: 3, col: 4}
	obj2 := Coordinates{row: 5, col: 5}
	a1 := Coordinates{row: 1, col: 3}
	a2 := Coordinates{row: 7, col: 6}
	r1 := calcAntinode(obj1, obj2)
	r2 := calcAntinode(obj2, obj1)
	if a1 != r1 {
		t.Error("A1 is different to R1")
	}
	if a2 != r2 {
		t.Error("A2 is different to R2")
	}
}
