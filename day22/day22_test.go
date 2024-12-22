package main

import (
	"teekometDev/filereader5"
	"testing"
)

func TestT1(t *testing.T) {
	file := filereader5.ReadFile("test.txt")
	got := Task1(file)
	want := 37327623

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestT2(t *testing.T) {
	file := filereader5.ReadFile("test2.txt")
	got := Task2(file)
	want := 23

	if got != want {
		t.Errorf("Got %d and wanted %d\n", got, want)
	}
}

func TestCreateSecretTimes1(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(1, &startInt, &secMap)
	want := 15887950

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}

func TestCreateSecretTimes2(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(2, &startInt, &secMap)
	want := 16495136

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}
func TestCreateSecretTimes3(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(3, &startInt, &secMap)
	want := 527345

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}
func TestCreateSecretTimes4(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(4, &startInt, &secMap)
	want := 704524

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}
func TestCreateSecretTimes5(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(5, &startInt, &secMap)
	want := 1553684

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}
func TestCreateSecretTimes6(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(6, &startInt, &secMap)
	want := 12683156

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}
func TestCreateSecretTimes7(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(7, &startInt, &secMap)
	want := 11100544

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}
func TestCreateSecretTimes8(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(8, &startInt, &secMap)
	want := 12249484

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}
func TestCreateSecretTimes9(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(9, &startInt, &secMap)
	want := 7753432

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}
func TestCreateSecretTimes10(t *testing.T) {
	secMap := make(map[int]int)
	startInt := 123
	calcSecret(10, &startInt, &secMap)
	want := 5908254

	if startInt != want {
		t.Errorf("Got %d and wanted %d\n", startInt, want)
	}
}

func TestDetermineSequences(t *testing.T) {
	secMap := make(map[int]int)
	gotMap := determineSequences(123, 10, &secMap)
	gotMax := 0
	wantMax := 6
	for _, entry := range gotMap {
		if entry > gotMax {
			gotMax = entry
		}
	}
	if gotMax != wantMax {
		t.Errorf("Got %d and wanted %d\n", gotMax, wantMax)
	}
}
