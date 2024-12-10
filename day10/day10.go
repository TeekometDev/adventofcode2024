package main

import (
	"fmt"
	"strconv"
	"teekometDev/filereader4"
	"time"
)

func main() {
	start := time.Now()
	file := filereader4.ReadFileAsMatrix("input.txt")
	res_1 := Task1(file)
	res_2 := Task2(file)
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res_1, res_2)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime: %v\n", duration.Microseconds())
}

func Task1(file [][]rune) int {
	bytemap := translateLandmap(&file)
	trailheads := findTrailheads(&bytemap)
	retNum := 0
	for _, trailhead := range trailheads {
		findPaths(&trailhead, &trailhead, &bytemap, false)
		retNum += evaluate(&trailhead)
	}
	return retNum
}

func Task2(file [][]rune) int {
	bytemap := translateLandmap(&file)
	trailheads := findTrailheads(&bytemap)
	retNum := 0
	for _, trailhead := range trailheads {
		findPaths(&trailhead, &trailhead, &bytemap, true)
		retNum += evaluate(&trailhead)
	}
	return retNum
}

func translateLandmap(landmap *[][]rune) [][]byte {
	var bytemap [][]byte
	for _, line := range *landmap {
		newLine := make([]byte, len((*landmap)[0]))
		for i, char := range line {
			newNum, _ := strconv.Atoi(string(char))
			newLine[i] = byte(newNum)
		}
		bytemap = append(bytemap, newLine)
	}
	return bytemap
}

func findTrailheads(bytemap *[][]byte) []Place {
	var places []Place
	for row_i, line := range *bytemap {
		for col_i, element := range line {
			if element == 0 {
				places = append(places, Place{row: row_i, col: col_i, height: 0})
			}
		}
	}
	return places
}

func findPaths(root *Place, element *Place, field *[][]byte, ignoreUniqueness bool) {
	height := len(*field)
	width := len((*field)[0])
	// Upwards
	nextRow, nextCol := (*element).row-1, (*element).col
	if checkPosition(nextRow, nextCol, (*element).height, field, height, width) {
		candidate := Place{row: nextRow, col: nextCol, height: (*field)[nextRow][nextCol]}
		if ignoreUniqueness || uniqueChild(root, &candidate) {
			(*element).children = append((*element).children, &candidate)
		}
	}
	// Downwards
	nextRow, nextCol = (*element).row+1, (*element).col
	if checkPosition(nextRow, nextCol, (*element).height, field, height, width) {
		candidate := Place{row: nextRow, col: nextCol, height: (*field)[nextRow][nextCol]}
		if ignoreUniqueness || uniqueChild(root, &candidate) {
			(*element).children = append((*element).children, &candidate)
		}
	}
	// Left
	nextRow, nextCol = (*element).row, (*element).col-1
	if checkPosition(nextRow, nextCol, (*element).height, field, height, width) {
		candidate := Place{row: nextRow, col: nextCol, height: (*field)[nextRow][nextCol]}
		if ignoreUniqueness || uniqueChild(root, &candidate) {
			(*element).children = append((*element).children, &candidate)
		}
	}
	// Right
	nextRow, nextCol = (*element).row, (*element).col+1
	if checkPosition(nextRow, nextCol, (*element).height, field, height, width) {
		candidate := Place{row: nextRow, col: nextCol, height: (*field)[nextRow][nextCol]}
		if ignoreUniqueness || uniqueChild(root, &candidate) {
			(*element).children = append((*element).children, &candidate)
		}
	}
	for _, vchild := range (*element).children {
		findPaths(root, vchild, field, ignoreUniqueness)
	}
}

func checkPosition(row int, col int, curLevel byte, field *[][]byte, height int, width int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= height || col >= width {
		return false
	}
	if curLevel != (*field)[row][col]-1 {
		return false
	}
	return true
}

func evaluate(place *Place) int {
	num := 0
	for _, child := range (*place).children {
		if child.height == 9 {
			num++
		} else {
			num += evaluate(child)
		}
	}
	return num
}
