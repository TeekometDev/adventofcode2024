package main

import (
	"fmt"
	"teekometDev/filereader5"
)

func main() {
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file)
	res2 := Task2(file)
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res1, res2)
}

func Task1(lines []string) int {
	var robot Robot
	box_coordinates := []Position{}
	obstacle_list := make(map[Position]bool)
	commands := false
	y_size := 0
	x_size := len(lines[0])
	for y, line := range lines {
		if line == "" {
			commands = true
			y_size = y
			continue
		}
		for x, char := range line {
			if commands {
				robot.movementList = append(robot.movementList, resolveCommand(char))
			} else if char == 'O' {
				box_coordinates = append(box_coordinates, Position{x_pos: x, y_pos: y})
			} else if char == '@' {
				robot = Robot{x_pos: x, y_pos: y}
			} else if char == '#' {
				obstacle_list[Position{x_pos: x, y_pos: y}] = true
			}
		}
	}
	playground := Warehouse{x_size: x_size, y_size: y_size}
	mapOfBoxes := createBoxMap(playground, &box_coordinates)
	movingRobot(&robot, &mapOfBoxes, &obstacle_list, playground, false)
	res := 0
	for coor, val := range mapOfBoxes {
		if val {
			res += 100*coor.y_pos + coor.x_pos
		}
	}
	return res
}

func movingRobot(robot *Robot, boxMap *map[Position]bool, obstacles *map[Position]bool, wareHouse Warehouse, printFld bool) {
	for _, move := range robot.movementList {
		moveRob(move, robot, boxMap, obstacles)
		if printFld {
			printField(wareHouse, robot, obstacles, boxMap)
		}
	}
}

func moveRob(move Movement, robot *Robot, boxMap *map[Position]bool, obstacles *map[Position]bool) {
	newx := robot.x_pos
	newy := robot.y_pos
	switch move {
	case Up:
		if checkBoundary(robot.x_pos, robot.y_pos-1, obstacles) {
			if checkBoxPos(robot.x_pos, robot.y_pos-1, boxMap) {
				if moveBoxesUp(robot.x_pos, robot.y_pos-1, boxMap, obstacles) {
					newy = robot.y_pos - 1
				}
			} else {
				newy = robot.y_pos - 1
			}
		}
	case Down:
		if checkBoundary(robot.x_pos, robot.y_pos+1, obstacles) {
			if checkBoxPos(robot.x_pos, robot.y_pos+1, boxMap) {
				if moveBoxesDown(robot.x_pos, robot.y_pos+1, boxMap, obstacles) {
					newy = robot.y_pos + 1
				}
			} else {
				newy = robot.y_pos + 1
			}
		}
	case Left:
		if checkBoundary(robot.x_pos-1, robot.y_pos, obstacles) {
			if checkBoxPos(robot.x_pos-1, robot.y_pos, boxMap) {
				if moveBoxesLeft(robot.x_pos-1, robot.y_pos, boxMap, obstacles) {
					newx = robot.x_pos - 1
				}
			} else {
				newx = robot.x_pos - 1
			}
		}
	case Right:
		if checkBoundary(robot.x_pos+1, robot.y_pos, obstacles) {
			if checkBoxPos(robot.x_pos+1, robot.y_pos, boxMap) {
				if moveBoxesRight(robot.x_pos+1, robot.y_pos, boxMap, obstacles) {
					newx = robot.x_pos + 1
				}
			} else {
				newx = robot.x_pos + 1
			}
		}
	}
	robot.x_pos = newx
	robot.y_pos = newy
}

func checkBoxPos(x_pos int, y_pos int, boxMap *map[Position]bool) bool {
	return (*boxMap)[Position{x_pos: x_pos, y_pos: y_pos}]
}

func checkBoundary(x_pos int, y_pos int, obstacles *map[Position]bool) bool {
	_, exists := (*obstacles)[Position{x_pos: x_pos, y_pos: y_pos}]
	return !exists
}

func moveBoxesUp(start_x int, start_y int, boxMap *map[Position]bool, obstacles *map[Position]bool) bool {
	curr_x := start_x
	curr_y := start_y
	isBox := true
	for isBox {
		curr_y = curr_y - 1
		if !checkBoundary(curr_x, curr_y, obstacles) {
			isBox = false
			return false
		}
		isBox = checkBoxPos(curr_x, curr_y, boxMap)
	}
	(*boxMap)[Position{x_pos: start_x, y_pos: start_y}] = false
	(*boxMap)[Position{x_pos: curr_x, y_pos: curr_y}] = true
	return true
}

func moveBoxesDown(start_x int, start_y int, boxMap *map[Position]bool, obstacles *map[Position]bool) bool {
	curr_x := start_x
	curr_y := start_y
	isBox := true
	for isBox {
		curr_y = curr_y + 1
		if !checkBoundary(curr_x, curr_y, obstacles) {
			isBox = false
			return false
		}
		isBox = checkBoxPos(curr_x, curr_y, boxMap)
	}
	(*boxMap)[Position{x_pos: start_x, y_pos: start_y}] = false
	(*boxMap)[Position{x_pos: curr_x, y_pos: curr_y}] = true
	return true
}

func moveBoxesLeft(start_x int, start_y int, boxMap *map[Position]bool, obstacles *map[Position]bool) bool {
	curr_x := start_x
	curr_y := start_y
	isBox := true
	for isBox {
		curr_x = curr_x - 1
		if !checkBoundary(curr_x, curr_y, obstacles) {
			isBox = false
			return false
		}
		isBox = checkBoxPos(curr_x, curr_y, boxMap)
	}
	(*boxMap)[Position{x_pos: start_x, y_pos: start_y}] = false
	(*boxMap)[Position{x_pos: curr_x, y_pos: curr_y}] = true
	return true
}

func moveBoxesRight(start_x int, start_y int, boxMap *map[Position]bool, obstacles *map[Position]bool) bool {
	curr_x := start_x
	curr_y := start_y
	isBox := true
	for isBox {
		curr_x = curr_x + 1
		if !checkBoundary(curr_x, curr_y, obstacles) {
			isBox = false
			return false
		}
		isBox = checkBoxPos(curr_x, curr_y, boxMap)
	}
	(*boxMap)[Position{x_pos: start_x, y_pos: start_y}] = false
	(*boxMap)[Position{x_pos: curr_x, y_pos: curr_y}] = true
	return true
}

func resolveCommand(char rune) Movement {
	switch char {
	case '^':
		return Up
	case '>':
		return Right
	case '<':
		return Left
	default:
		return Down
	}
}

func createBoxMap(warehouse Warehouse, boxes *[]Position) map[Position]bool {
	boxmap := make(map[Position]bool)
	for x := range warehouse.x_size {
		for y := range warehouse.y_size {
			boxmap[Position{x_pos: x, y_pos: y}] = false
		}
	}
	for _, pos := range *boxes {
		boxmap[pos] = true
	}
	return boxmap
}
