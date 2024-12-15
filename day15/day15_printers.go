package main

import "fmt"

func printField(playground Warehouse, robot *Robot, obstacles *map[Position]bool, boxes *map[Position]bool) {
	for y := range playground.y_size {
		line := ""
		for x := range playground.x_size {
			currPos := Position{x_pos: x, y_pos: y}
			_, obstacleExists := (*obstacles)[currPos]
			if y == robot.y_pos && x == robot.x_pos {
				line += "@"
			} else if obstacleExists {
				line += "#"
			} else if (*boxes)[currPos] {
				line += "O"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func printDoubledPlayground(warehouse *[][]FieldObject, robot *Robot) {
	for y, row := range *warehouse {
		line := ""
		for x, char := range row {
			if x == robot.x_pos && y == robot.y_pos {
				line += "@"
			} else {
				switch char {
				case Empty:
					line += "."
				case Boundary:
					line += "#"
				case BoxLeft:
					line += "["
				case BoxRight:
					line += "]"
				}
			}
		}
		fmt.Println(line)
	}
}
