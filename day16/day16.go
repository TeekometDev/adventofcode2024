package main

import (
	"fmt"
	"teekometDev/filereader4"
)

func main() {
	input := filereader4.ReadFileAsMatrix("input.txt")
	res1 := Task1(input)
	res2 := 0
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res1, res2)
}

func Task1(file [][]rune) int {
	playground, start_x, start_y := initPlayground(file)
	dijkstra(&playground, start_x, start_y)
	printField(&playground)
	for _, line := range playground {
		for _, element := range line {
			if element.isTarget {
				return element.distance
			}
		}
	}
	return -1
}

func Task2(file [][]rune) int {
	return 0
}

func initPlayground(input [][]rune) ([][]Field, int, int) {
	start_x := -1
	start_y := -1
	retField := make([][]Field, len(input))
	for y, line := range input {
		newLine := make([]Field, len(line))
		for x, char := range line {
			if char == '#' {
				newLine[x] = Field{isTarget: false, isBoundary: true}
			}
			if char == '.' {
				newLine[x] = Field{isTarget: false, isBoundary: false, distance: -1}
			}
			if char == 'E' {
				newLine[x] = Field{isTarget: true, isBoundary: false, distance: -1}
			}
			if char == 'S' {
				start_x = x
				start_y = y
				newLine[x] = Field{isTarget: false, isBoundary: false, distance: 0, direction: Right}
			}
		}
		retField[y] = newLine
	}
	return retField, start_x, start_y
}

func dijkstra(field *[][]Field, start_x int, start_y int) {
	discoverList := [][]int{{start_x, start_y}}
	for len(discoverList) > 0 {
		currentElement := discoverList[0]
		fieldElement := (*field)[currentElement[1]][currentElement[0]]
		discoverList = discoverList[1:]
		// Check Upwards
		upElement := (*field)[currentElement[1]-1][currentElement[0]]
		if newElement(upElement) {
			tempDistance := fieldElement.distance + 1
			tempDistance += 1000 * turning(fieldElement.direction, Up)
			if upElement.distance < 0 || tempDistance < upElement.distance {
				(*field)[currentElement[1]-1][currentElement[0]].direction = Up
				(*field)[currentElement[1]-1][currentElement[0]].distance = tempDistance
				discoverList = append(discoverList, []int{currentElement[0], currentElement[1] - 1})
			}
		}
		// Check Right
		rightElement := (*field)[currentElement[1]][currentElement[0]+1]
		if newElement(rightElement) {
			tempDistance := fieldElement.distance + 1
			tempDistance += 1000 * turning(fieldElement.direction, Right)
			if rightElement.distance < 0 || tempDistance < rightElement.distance {
				(*field)[currentElement[1]][currentElement[0]+1].direction = Right
				(*field)[currentElement[1]][currentElement[0]+1].distance = tempDistance
				discoverList = append(discoverList, []int{currentElement[0] + 1, currentElement[1]})
			}
		}
		// Check Downwards
		downElement := (*field)[currentElement[1]+1][currentElement[0]]
		if newElement(downElement) {
			tempDistance := fieldElement.distance + 1
			tempDistance += 1000 * turning(fieldElement.direction, Down)
			if downElement.distance < 0 || tempDistance < downElement.distance {
				(*field)[currentElement[1]+1][currentElement[0]].direction = Down
				(*field)[currentElement[1]+1][currentElement[0]].distance = tempDistance
				discoverList = append(discoverList, []int{currentElement[0], currentElement[1] + 1})
			}
		}
		// Check Left
		leftElement := (*field)[currentElement[1]][currentElement[0]-1]
		if newElement(leftElement) {
			tempDistance := fieldElement.distance + 1
			tempDistance += 1000 * turning(fieldElement.direction, Left)
			if leftElement.distance < 0 || tempDistance < leftElement.distance {
				(*field)[currentElement[1]][currentElement[0]-1].direction = Left
				(*field)[currentElement[1]][currentElement[0]-1].distance = tempDistance
				discoverList = append(discoverList, []int{currentElement[0] - 1, currentElement[1]})
			}
		}
	}
}

func newElement(el Field) bool {
	return !el.isBoundary
}

func turning(oldDirection Direction, newDirection Direction) int {
	if oldDirection > newDirection {
		return int(oldDirection) - int(newDirection)
	}
	return int(newDirection) - int(oldDirection)
}

func printField(field *[][]Field) {
	for _, line := range *field {
		newLine := ""
		for _, char := range line {
			if char.isBoundary {
				newLine += "#"
				continue
			}
			switch char.direction {
			case Up:
				newLine += "^"
			case Right:
				newLine += ">"
			case Down:
				newLine += "v"
			case Left:
				newLine += "<"
			}
		}
		fmt.Println(newLine)
	}
}
