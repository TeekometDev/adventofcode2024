package main

import (
	"fmt"
	"strconv"
	"strings"
	"teekometDev/filereader5"
)

func main() {
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file, 70, 70, 1024)
	res2x, res2y := Task2(file, 70, 70)
	fmt.Printf("RESULT 1: %d, RESULT 2: %d, %d\n", res1, res2x, res2y)
}

func Task1(file []string, xDim int, yDim int, bytesToFall int) int {
	var field [][]int
	for range yDim + 1 {
		var line []int
		for range xDim + 1 {
			line = append(line, -1)
		}
		field = append(field, line)
	}
	field[0][0] = 0
	obstacles := makeObstacles(file, bytesToFall)
	dijkstra(&field, &obstacles)
	return field[yDim][xDim]
}

func Task2(file []string, xDim int, yDim int) (int, int) {
	bytes := 0
	retVal := 0
	for retVal > -1 {
		bytes += 1
		retVal = Task1(file, xDim, yDim, bytes)
	}
	strResult := strings.Split(file[bytes-1], ",")
	xcoor, _ := strconv.Atoi(strResult[0])
	ycoor, _ := strconv.Atoi(strResult[1])
	return xcoor, ycoor
}

func dijkstra(field *[][]int, obstacles *map[Coordinate]bool) {
	toCheck := []Coordinate{{x: 0, y: 0}}
	for len(toCheck) > 0 {
		element := toCheck[0]
		toCheck = toCheck[1:]
		tempVal := (*field)[element.y][element.x] + 1
		if relevantField(element.x, element.y-1, field, obstacles) {
			newx := element.x
			newy := element.y - 1
			fieldVal := (*field)[newy][newx]
			if fieldVal == -1 || tempVal < fieldVal {
				(*field)[newy][newx] = tempVal
				toCheck = append(toCheck, Coordinate{x: newx, y: newy})
			}
		}
		if relevantField(element.x, element.y+1, field, obstacles) {
			newx := element.x
			newy := element.y + 1
			fieldVal := (*field)[newy][newx]
			if fieldVal == -1 || tempVal < fieldVal {
				(*field)[newy][newx] = tempVal
				toCheck = append(toCheck, Coordinate{x: newx, y: newy})
			}
		}
		if relevantField(element.x-1, element.y, field, obstacles) {
			newx := element.x - 1
			newy := element.y
			fieldVal := (*field)[newy][newx]
			if fieldVal == -1 || tempVal < fieldVal {
				(*field)[newy][newx] = tempVal
				toCheck = append(toCheck, Coordinate{x: newx, y: newy})
			}
		}
		if relevantField(element.x+1, element.y, field, obstacles) {
			newx := element.x + 1
			newy := element.y
			fieldVal := (*field)[newy][newx]
			if fieldVal == -1 || tempVal < fieldVal {
				(*field)[newy][newx] = tempVal
				toCheck = append(toCheck, Coordinate{x: newx, y: newy})
			}
		}
	}
}

func relevantField(x_coor int, y_coor int, field *[][]int, obstacles *map[Coordinate]bool) bool {
	if x_coor >= len((*field)[0]) || x_coor < 0 {
		return false
	}
	if y_coor >= len((*field)) || y_coor < 0 {
		return false
	}
	_, exists := (*obstacles)[Coordinate{x: x_coor, y: y_coor}]
	return !exists
}

func makeObstacles(file []string, bytes int) map[Coordinate]bool {
	obs := make(map[Coordinate]bool)
	for i := range bytes {
		lineArr := strings.Split(file[i], ",")
		x_coor, _ := strconv.Atoi(lineArr[0])
		y_coor, _ := strconv.Atoi(lineArr[1])
		obs[Coordinate{x: x_coor, y: y_coor}] = true
	}
	return obs
}
