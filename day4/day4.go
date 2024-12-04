package main

import (
	"fmt"
	"regexp"
	"teekometDev/filereader4"
	"teekometDev/matrixhelpers"
)

func Task1(fileName string) int {
	counter := 0
	horiMatr := filereader4.ReadFileAsMatrix(fileName)
	vertMatr := matrixhelpers.CreateVertical(horiMatr)
	turnMatr_h := matrixhelpers.RotateMatrix45(horiMatr)
	turnMatr_v := matrixhelpers.RotateMatrix135(vertMatr)
	counter += evaluateMatrix(horiMatr)
	counter += evaluateMatrix(vertMatr)
	counter += evaluateMatrix(turnMatr_h)
	counter += evaluateMatrix(turnMatr_v)
	return counter
}

func Task2(fileName string) int {
	counter := 0
	matr := filereader4.ReadFileAsMatrix(fileName)
	rows := len(matr)
	cols := len(matr[0])
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if checkCrossing(string([]rune{matr[i-1][j-1], matr[i][j], matr[i+1][j+1]}), string([]rune{matr[i+1][j-1], matr[i][j], matr[i-1][j+1]})) {
				counter++
			}
		}
	}
	return counter
}

func evaluateMatrix(input [][]rune) int {
	result := 0
	for _, line := range input {
		result += calculateLine(line)
	}
	return result
}

func calculateLine(input []rune) int {
	result := 0
	r, _ := regexp.Compile(`XMAS|SAMX`)
	for i := 0; i < len(input)-3; i++ {
		newStr := string(input[i : i+4])
		if r.FindString(newStr) != "" {
			result++
		}
	}
	return result
}

func checkCrossing(word1 string, word2 string) bool {
	r, _ := regexp.Compile(`MAS|SAM`)
	is1 := r.FindString(word1) != ""
	is2 := r.FindString(word2) != ""
	return is1 && is2
}

func main() {
	result1 := Task1("input.txt")
	result2 := Task2("input.txt")
	fmt.Printf("Result 1: %d\nResult 2: %d\n", result1, result2)
}
