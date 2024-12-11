package main

import "testing"

func TestT1(t *testing.T) {
	got := Task1("125 17")
	want := 55312

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}
