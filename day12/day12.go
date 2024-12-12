package main

import (
	"fmt"
	"teekometDev/filereader4"
)

func main() {
	file := filereader4.ReadFileAsMatrix("input.txt")
	res_1 := Task1(file)
	res_2 := Task2(filereader4.ReadFileAsMatrix("input.txt"))
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res_1, res_2)
}

func Task1(input [][]rune) int {
	retVal := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '.' {
				retVal += solveChar(i, j, &input)
			}
		}
	}
	return retVal
}

func Task2(input [][]rune) int {
	retVal := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '.' {
				retVal += solveChar2(i, j, &input)
			}
		}
	}
	return retVal
}

func solveChar(row int, col int, field *[][]rune) int {
	listToSolve := []Coordinate{{row: row, col: col}}
	listToRemove := make(map[Coordinate]bool)
	listToRemove[Coordinate{row: row, col: col}] = true
	currChar := (*field)[row][col]
	fences := 0
	area := 0
	for len(listToSolve) > 0 {
		area++
		coordinate := listToSolve[0]
		listToSolve = listToSolve[1:]
		fences += checkUpwards(coordinate, currChar, field, &listToSolve, &listToRemove)
		fences += checkDownwards(coordinate, currChar, field, &listToSolve, &listToRemove)
		fences += checkLeft(coordinate, currChar, field, &listToSolve, &listToRemove)
		fences += checkRight(coordinate, currChar, field, &listToSolve, &listToRemove)
	}
	for key := range listToRemove {
		(*field)[key.row][key.col] = '.'
	}
	return fences * area
}

func solveChar2(row int, col int, field *[][]rune) int {
	listToSolve := []Coordinate{{row: row, col: col}}
	listToRemove := make(map[Coordinate]bool)
	listToRemove[Coordinate{row: row, col: col}] = true
	listOfFences := make(map[Coordinate]Fences)
	currChar := (*field)[row][col]
	area := 0
	for len(listToSolve) > 0 {
		area++
		coordinate := listToSolve[0]
		listToSolve = listToSolve[1:]
		listOfFences[coordinate] = Fences{}
		if checkUpwards(coordinate, currChar, field, &listToSolve, &listToRemove) == 1 {
			fence := listOfFences[coordinate]
			fence.up = true
			listOfFences[coordinate] = fence
		}
		if checkDownwards(coordinate, currChar, field, &listToSolve, &listToRemove) == 1 {
			fence := listOfFences[coordinate]
			fence.down = true
			listOfFences[coordinate] = fence
		}
		if checkLeft(coordinate, currChar, field, &listToSolve, &listToRemove) == 1 {
			fence := listOfFences[coordinate]
			fence.left = true
			listOfFences[coordinate] = fence
		}
		if checkRight(coordinate, currChar, field, &listToSolve, &listToRemove) == 1 {
			fence := listOfFences[coordinate]
			fence.right = true
			listOfFences[coordinate] = fence
		}
	}
	fences := calcFenceNr(&listOfFences)
	for key := range listToRemove {
		(*field)[key.row][key.col] = '.'
	}
	return fences * area
}

func checkUpwards(coor Coordinate, character rune, field *[][]rune, solveList *[]Coordinate, checkRef *map[Coordinate]bool) int {
	if coor.row == 0 {
		return 1
	}
	if check(coor.row-1, coor.col, character, field) {
		newCoordinate := Coordinate{row: coor.row - 1, col: coor.col}
		if _, exists := (*checkRef)[newCoordinate]; !exists {
			(*solveList) = append((*solveList), newCoordinate)
			(*checkRef)[newCoordinate] = true
		}
		return 0
	} else {
		return 1
	}
}

func checkDownwards(coor Coordinate, character rune, field *[][]rune, solveList *[]Coordinate, checkRef *map[Coordinate]bool) int {
	if coor.row == len(*field)-1 {
		return 1
	}
	if check(coor.row+1, coor.col, character, field) {
		newCoordinate := Coordinate{row: coor.row + 1, col: coor.col}
		if _, exists := (*checkRef)[newCoordinate]; !exists {
			(*solveList) = append((*solveList), newCoordinate)
			(*checkRef)[newCoordinate] = true
		}
		return 0
	} else {
		return 1
	}
}

func checkLeft(coor Coordinate, character rune, field *[][]rune, solveList *[]Coordinate, checkRef *map[Coordinate]bool) int {
	if coor.col == 0 {
		return 1
	}
	if check(coor.row, coor.col-1, character, field) {
		newCoordinate := Coordinate{row: coor.row, col: coor.col - 1}
		if _, exists := (*checkRef)[newCoordinate]; !exists {
			(*solveList) = append((*solveList), newCoordinate)
			(*checkRef)[newCoordinate] = true
		}
		return 0
	} else {
		return 1
	}
}

func checkRight(coor Coordinate, character rune, field *[][]rune, solveList *[]Coordinate, checkRef *map[Coordinate]bool) int {
	if coor.col == len((*field)[0])-1 {
		return 1
	}
	if check(coor.row, coor.col+1, character, field) {
		newCoordinate := Coordinate{row: coor.row, col: coor.col + 1}
		if _, exists := (*checkRef)[newCoordinate]; !exists {
			(*solveList) = append((*solveList), newCoordinate)
			(*checkRef)[newCoordinate] = true
		}
		return 0
	} else {
		return 1
	}
}

func check(row int, col int, character rune, field *[][]rune) bool {
	return (*field)[row][col] == character
}

func calcFenceNr(fences *map[Coordinate]Fences) int {
	fenceNr := 0
	for coor := range *fences {
		fence := (*fences)[coor]
		if fence.up {
			fenceNr++
			checkFenceLeftwards(coor, true, false, fences)
			checkFenceRightwards(coor, true, false, fences)
		}
		if fence.down {
			fenceNr++
			checkFenceLeftwards(coor, false, true, fences)
			checkFenceRightwards(coor, false, true, fences)
		}
		if fence.left {
			fenceNr++
			checkFenceUpwards(coor, true, false, fences)
			checkFenceDownwards(coor, true, false, fences)
		}
		if fence.right {
			fenceNr++
			checkFenceUpwards(coor, false, true, fences)
			checkFenceDownwards(coor, false, true, fences)
		}
	}
	return fenceNr
}

func checkFenceLeftwards(coordinate Coordinate, up bool, down bool, fenceMap *map[Coordinate]Fences) {
	update := (*fenceMap)[coordinate]
	next := Coordinate{row: coordinate.row, col: coordinate.col - 1}
	newFences, exists := (*fenceMap)[next]
	if up {
		update.up = false
		(*fenceMap)[coordinate] = update
		if exists && newFences.up {
			checkFenceLeftwards(next, up, down, fenceMap)
		}
	} else if down {
		update.down = false
		(*fenceMap)[coordinate] = update
		if exists && newFences.down {
			checkFenceLeftwards(next, up, down, fenceMap)
		}
	}
}

func checkFenceRightwards(coordinate Coordinate, up bool, down bool, fenceMap *map[Coordinate]Fences) {
	update := (*fenceMap)[coordinate]
	next := Coordinate{row: coordinate.row, col: coordinate.col + 1}
	newFences, exists := (*fenceMap)[next]
	if up {
		update.up = false
		(*fenceMap)[coordinate] = update
		if exists && newFences.up {
			checkFenceRightwards(next, up, down, fenceMap)
		}
	} else if down {
		update.down = false
		(*fenceMap)[coordinate] = update
		if exists && newFences.down {
			checkFenceRightwards(next, up, down, fenceMap)
		}
	}
}

func checkFenceUpwards(coordinate Coordinate, left bool, right bool, fenceMap *map[Coordinate]Fences) {
	update := (*fenceMap)[coordinate]
	next := Coordinate{row: coordinate.row - 1, col: coordinate.col}
	newFences, exists := (*fenceMap)[next]
	if right {
		update.right = false
		(*fenceMap)[coordinate] = update
		if exists && newFences.right {
			checkFenceUpwards(next, left, right, fenceMap)
		}
	} else if left {
		update.left = false
		(*fenceMap)[coordinate] = update
		if exists && newFences.left {
			checkFenceUpwards(next, left, right, fenceMap)
		}
	}
}

func checkFenceDownwards(coordinate Coordinate, left bool, right bool, fenceMap *map[Coordinate]Fences) {
	update := (*fenceMap)[coordinate]
	next := Coordinate{row: coordinate.row + 1, col: coordinate.col}
	newFences, exists := (*fenceMap)[next]
	if right {
		update.right = false
		(*fenceMap)[coordinate] = update
		if exists && newFences.right {
			checkFenceDownwards(next, left, right, fenceMap)
		}
	} else if left {
		update.left = false
		(*fenceMap)[coordinate] = update
		if exists && newFences.left {
			checkFenceDownwards(next, left, right, fenceMap)
		}
	}
}
