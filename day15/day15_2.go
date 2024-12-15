package main

import "fmt"

func Task2(file []string) int {
	var robot Robot
	var warehouse [][]FieldObject
	commands := false
	for y, line := range file {
		if line == "" {
			commands = true
			continue
		}
		var warehouseLine []FieldObject
		for x, char := range line {
			if commands {
				robot.movementList = append(robot.movementList, resolveCommand(char))
			} else if char == 'O' {
				warehouseLine = append(warehouseLine, BoxLeft)
				warehouseLine = append(warehouseLine, BoxRight)
			} else if char == '@' {
				robot = Robot{x_pos: x * 2, y_pos: y}
				warehouseLine = append(warehouseLine, Empty)
				warehouseLine = append(warehouseLine, Empty)
			} else if char == '#' {
				warehouseLine = append(warehouseLine, Boundary)
				warehouseLine = append(warehouseLine, Boundary)
			} else if char == '.' {
				warehouseLine = append(warehouseLine, Empty)
				warehouseLine = append(warehouseLine, Empty)
			}
		}
		warehouse = append(warehouse, warehouseLine)
	}
	fmt.Println("Initial:")
	for _, move := range robot.movementList {
		printDoubledPlayground(&warehouse, &robot)
		fmt.Printf("Next Move: %d\n", move)
		moveRobot2(&robot, &warehouse, move)
	}
	printDoubledPlayground(&warehouse, &robot)
	return 0
}

func moveRobot2(robot *Robot, playground *[][]FieldObject, move Movement) {
	x_pos := robot.x_pos
	y_pos := robot.y_pos
	switch move {
	case Up:
		y_pos -= 1
		if !moveIfPossible(robot, x_pos, y_pos, playground) {
			fmt.Printf("New Position: %d %d\n", y_pos, x_pos)
			if checkBoxesUp(x_pos, y_pos, playground, true) {
				moveBigBoxesUp(x_pos, y_pos, playground)
				robot.y_pos = y_pos
			}
		}
	case Down:
		y_pos += 1
		if !moveIfPossible(robot, x_pos, y_pos, playground) {
			if checkBoxesDown(x_pos, y_pos, playground, true) {
				moveBigBoxesDown(x_pos, y_pos, playground, true)
				robot.y_pos = y_pos
			}
		}
	case Left:
		x_pos -= 1
		if !moveIfPossible(robot, x_pos, y_pos, playground) {
			if checkBoxesLeft(x_pos, y_pos, playground) {
				moveBigBoxesLeft(x_pos, y_pos, playground)
				robot.x_pos = x_pos
			}
		}
	case Right:
		x_pos += 1
		if !moveIfPossible(robot, x_pos, y_pos, playground) {
			if checkBoxesRight(x_pos, y_pos, playground) {
				moveBigBoxesRight(x_pos, y_pos, playground)
				robot.x_pos = x_pos
			}
		}
	}
}

func moveIfPossible(robot *Robot, new_x int, new_y int, field *[][]FieldObject) bool {
	char := (*field)[new_y][new_x]
	switch char {
	case Empty:
		robot.x_pos = new_x
		robot.y_pos = new_y
		return true
	case Boundary:
		return true
	default:
		return false
	}
}

func checkBoxesUp(x_pos int, y_pos int, field *[][]FieldObject, first bool) bool {
	if first {
		if (*field)[y_pos][x_pos] == BoxLeft {
			fmt.Printf("Checking: %d %d and %d %d\n", y_pos, x_pos, y_pos, x_pos+1)
			return checkBoxesUp(x_pos, y_pos, field, false) && checkBoxesUp(x_pos+1, y_pos, field, false)
		} else {
			fmt.Printf("Checking: %d %d and %d %d\n", y_pos, x_pos, y_pos, x_pos-1)
			return checkBoxesUp(x_pos-1, y_pos, field, false) && checkBoxesUp(x_pos, y_pos, field, false)
		}
	}
	currElement := (*field)[y_pos][x_pos]
	if currElement == Boundary {
		return false
	}
	if currElement == BoxLeft {
		return checkBoxesUp(x_pos, y_pos-1, field, false) && checkBoxesUp(x_pos+1, y_pos-1, field, false)
	}
	if currElement == BoxRight {
		return checkBoxesUp(x_pos-1, y_pos-1, field, false) && checkBoxesUp(x_pos, y_pos-1, field, false)
	}
	return true
}

func moveBigBoxesUp(x_pos int, y_pos int, field *[][]FieldObject) {
	currElement := (*field)[y_pos][x_pos]
	if currElement == BoxLeft {
		moveBigBoxesUp(x_pos, y_pos-1, field)
		moveBigBoxesUp(x_pos+1, y_pos-1, field)
		(*field)[y_pos][x_pos] = Empty
		(*field)[y_pos-1][x_pos] = currElement
	}
	if currElement == BoxRight {
		moveBigBoxesUp(x_pos-1, y_pos-1, field)
		moveBigBoxesUp(x_pos, y_pos-1, field)
		(*field)[y_pos][x_pos] = Empty
		(*field)[y_pos-1][x_pos] = currElement
	}
}

func checkBoxesDown(x_pos int, y_pos int, field *[][]FieldObject, first bool) bool {
	if first {
		if (*field)[y_pos][x_pos] == BoxLeft {
			return checkBoxesDown(x_pos, y_pos, field, false) && checkBoxesDown(x_pos+1, y_pos, field, false)
		} else {
			return checkBoxesDown(x_pos-1, y_pos, field, false) && checkBoxesDown(x_pos, y_pos, field, false)
		}
	}
	currElement := (*field)[y_pos][x_pos]
	if currElement == Boundary {
		return false
	}
	if currElement == BoxLeft {
		return checkBoxesDown(x_pos, y_pos+1, field, false) && checkBoxesDown(x_pos+1, y_pos+1, field, false)
	}
	if currElement == BoxRight {
		return checkBoxesDown(x_pos-1, y_pos+1, field, false) && checkBoxesDown(x_pos, y_pos+1, field, false)
	}
	return true
}

func moveBigBoxesDown(x_pos int, y_pos int, field *[][]FieldObject, first bool) {
	if first {
		if (*field)[y_pos][x_pos] == BoxLeft {
			moveBigBoxesDown(x_pos, y_pos, field, false)
			moveBigBoxesDown(x_pos+1, y_pos, field, false)
			return
		} else {
			moveBigBoxesDown(x_pos-1, y_pos, field, false)
			moveBigBoxesDown(x_pos, y_pos, field, false)
			return
		}
	}
	currElement := (*field)[y_pos][x_pos]
	if currElement == BoxLeft {
		moveBigBoxesDown(x_pos, y_pos+1, field, false)
		moveBigBoxesDown(x_pos+1, y_pos+1, field, false)
		(*field)[y_pos][x_pos] = Empty
		(*field)[y_pos+1][x_pos] = currElement
	}
	if currElement == BoxRight {
		moveBigBoxesDown(x_pos-1, y_pos+1, field, false)
		moveBigBoxesDown(x_pos, y_pos+1, field, false)
		(*field)[y_pos][x_pos] = Empty
		(*field)[y_pos+1][x_pos] = currElement
	}

}

func checkBoxesLeft(x_pos int, y_pos int, field *[][]FieldObject) bool {
	currElement := (*field)[y_pos][x_pos]
	if currElement == Boundary {
		return false
	}
	if currElement == BoxLeft || currElement == BoxRight {
		return checkBoxesLeft(x_pos-1, y_pos, field)
	}
	return true
}

func moveBigBoxesLeft(x_pos int, y_pos int, field *[][]FieldObject) {
	currElement := (*field)[y_pos][x_pos]
	if currElement == BoxLeft || currElement == BoxRight {
		moveBigBoxesLeft(x_pos-1, y_pos, field)
		(*field)[y_pos][x_pos] = Empty
		(*field)[y_pos][x_pos-1] = currElement
	}
}

func checkBoxesRight(x_pos int, y_pos int, field *[][]FieldObject) bool {
	currElement := (*field)[y_pos][x_pos]
	if currElement == Boundary {
		return false
	}
	if currElement == BoxLeft || currElement == BoxRight {
		return checkBoxesRight(x_pos+1, y_pos, field)
	}
	return true
}

func moveBigBoxesRight(x_pos int, y_pos int, field *[][]FieldObject) {
	currElement := (*field)[y_pos][x_pos]
	if currElement == BoxLeft || currElement == BoxRight {
		moveBigBoxesRight(x_pos+1, y_pos, field)
		(*field)[y_pos][x_pos] = Empty
		(*field)[y_pos][x_pos+1] = currElement
	}
}
