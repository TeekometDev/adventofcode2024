package main

import (
	"fmt"
	"teekometDev/filereader4"
)

func main() {
	result1 := Task1(getFile("input.txt"))
	result2 := Task2(getFile("input.txt"))
	fmt.Printf("RESULT 1: %d, RESULT 2: %d", result1, result2)
}

func getFile(fileName string) [][]rune {
	return filereader4.ReadFileAsMatrix(fileName)
}

func Task1(file [][]rune) int {
	guard := findStart(&file)
	file[guard.row][guard.col] = 'X'
	result := 1
	inField := true
	for inField {
		inField = move(&guard, &file)
		if file[guard.row][guard.col] != 'X' {
			result++
			file[guard.row][guard.col] = 'X'
		}
	}
	return result
}

func Task2(file [][]rune) int {
	guard := findStart(&file)
	// Create the valid path
	start_guard := guard
	listOfVisits := []Guard{start_guard}
	inField := true
	for inField {
		inField = move(&guard, &file)
		listOfVisits = append(listOfVisits, guard)
	}
	result := 0
	alreadyTested := make([][]bool, len(file))
	for i := range alreadyTested {
		alreadyTested[i] = make([]bool, len(file[0]))
	}
	for i := 1; i < len(listOfVisits); i++ {
		if createLoopObstacle(listOfVisits[0], listOfVisits[i], file, &alreadyTested) {
			result++
		}
	}
	return result
}

func findStart(grid *[][]rune) Guard {
	for i, line := range *grid {
		for j, char := range line {
			if string(char) == "^" {
				return Guard{col: j, row: i, facing: Upwards}
			}
		}
	}
	return Guard{}
}

func move(guard *Guard, grid *[][]rune) bool {
	rows := len(*grid)
	cols := len((*grid)[0])
	switch guard.facing {
	case Upwards:
		if !inBoundaries(guard.row-1, rows) {
			return false
		}
		if isClear((*grid)[guard.row-1][guard.col]) {
			guard.row = guard.row - 1
		} else {
			guard.facing = Right
			return move(guard, grid)
		}
	case Downwards:
		if !inBoundaries(guard.row+1, rows) {
			return false
		}
		if isClear((*grid)[guard.row+1][guard.col]) {
			guard.row = guard.row + 1
		} else {
			guard.facing = Left
			return move(guard, grid)
		}
	case Left:
		if !inBoundaries(guard.col-1, cols) {
			return false
		}
		if isClear((*grid)[guard.row][guard.col-1]) {
			guard.col = guard.col - 1
		} else {
			guard.facing = Upwards
			return move(guard, grid)
		}
	case Right:
		if !inBoundaries(guard.col+1, cols) {
			return false
		}
		if isClear((*grid)[guard.row][guard.col+1]) {
			guard.col = guard.col + 1
		} else {
			guard.facing = Downwards
			return move(guard, grid)
		}
	}
	return true
}

func inBoundaries(new_index int, size int) bool {
	if new_index < 0 {
		return false
	}
	if new_index >= size {
		return false
	}
	return true
}

func isClear(char rune) bool {
	return char != '#' && char != 'O'
}

func createLoopObstacle(start Guard, next Guard, room [][]rune, tested *[][]bool) bool {
	if (*tested)[next.row][next.col] {
		return false
	} else {
		(*tested)[next.row][next.col] = true
	}
	room[next.row][next.col] = 'O'
	inField := true
	fieldList := makeMatrix(len(room), len(room[0]))
	setDirection(&start, &fieldList[start.row][start.col])
	for inField {
		inField = move2(&start, &room)
		if inField && setDirection(&start, &fieldList[start.row][start.col]) {
			room[next.row][next.col] = '.'
			return true
		}
	}
	room[next.row][next.col] = '.'
	return false
}

func makeMatrix(rows int, cols int) [][]DirectionTile {
	matr := make([][]DirectionTile, rows)
	for i := range matr {
		matr[i] = make([]DirectionTile, cols)
	}
	return matr
}

func setDirection(guard *Guard, actCell *DirectionTile) bool {
	switch guard.facing {
	case Upwards:
		if actCell.up {
			return true
		}
		actCell.up = true
	case Downwards:
		if actCell.down {
			return true
		}
		actCell.down = true
	case Left:
		if actCell.left {
			return true
		}
		actCell.left = true
	case Right:
		if actCell.right {
			return true
		}
		actCell.right = true
	}
	return false
}

// Without recursive function calls
func move2(guard *Guard, grid *[][]rune) bool {
	rows := len(*grid)
	cols := len((*grid)[0])
	switch guard.facing {
	case Upwards:
		if !inBoundaries(guard.row-1, rows) {
			return false
		}
		if isClear((*grid)[guard.row-1][guard.col]) {
			guard.row = guard.row - 1
		} else {
			guard.facing = Right
			return true
		}
	case Downwards:
		if !inBoundaries(guard.row+1, rows) {
			return false
		}
		if isClear((*grid)[guard.row+1][guard.col]) {
			guard.row = guard.row + 1
		} else {
			guard.facing = Left
			return true
		}
	case Left:
		if !inBoundaries(guard.col-1, cols) {
			return false
		}
		if isClear((*grid)[guard.row][guard.col-1]) {
			guard.col = guard.col - 1
		} else {
			guard.facing = Upwards
			return true
		}
	case Right:
		if !inBoundaries(guard.col+1, cols) {
			return false
		}
		if isClear((*grid)[guard.row][guard.col+1]) {
			guard.col = guard.col + 1
		} else {
			guard.facing = Downwards
			return true
		}
	}
	return true
}

func printMatrix(place [][]rune) {
	for _, line := range place {
		fmt.Println(string(line))
	}
}
