package main

import "testing"

func Test1(t *testing.T) {
	got := Task1("test.txt")
	want := 143

	if got != want {
		t.Errorf("Got %d instead of %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := Task2("test.txt")
	want := 123

	if got != want {
		t.Errorf("Got %d instead of %d", got, want)
	}
}

func TestCalc1(t *testing.T) {
	input := [][]int{{75, 47, 61, 53, 29}, {97, 61, 53, 29, 13}, {75, 29, 13}}
	got := calcResult1(input)
	want := 143

	if got != want {
		t.Errorf("Got %d instead of %d", got, want)
	}
}

func TestT1Result(t *testing.T) {
	got := Task1("input.txt")
	want := 5588

	if got != want {
		t.Errorf("Got %d instead of %d", got, want)
	}
}
