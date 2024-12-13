package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"teekometDev/filereader5"

	"gonum.org/v1/gonum/mat"
)

func main() {
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file)
	res2 := Task2(file)
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res1, res2)
}

func Task1(input []string) int {
	var button1 Button
	var button2 Button
	retVal := 0
	for i := 0; i < len(input); i++ {
		line := input[i]
		if i%4 == 0 {
			button1 = createButton(line)
			button1.costs = 3
		}
		if i%4 == 1 {
			button2 = createButton(line)
			button2.costs = 1
		}
		if i%4 == 2 {
			fmt.Println(button1)
			xTar, yTar := createTarget(line)
			retVal += calculateValues(button1, button2, xTar, yTar)
		}
	}
	return retVal
}

func Task2(input []string) int {
	var button1 Button
	var button2 Button
	retVal := 0
	for i := 0; i < len(input); i++ {
		line := input[i]
		if i%4 == 0 {
			button1 = createButton(line)
			button1.costs = 3
		}
		if i%4 == 1 {
			button2 = createButton(line)
			button2.costs = 1
		}
		if i%4 == 2 {
			fmt.Println(button1)
			xTar, yTar := createTarget(line)
			retVal += calculateValues2(button1, button2, xTar+10000000000000, yTar+10000000000000)
		}
	}
	return retVal
}

func calculateValues(button1 Button, button2 Button, xTarget int, yTarget int) int {
	matrA := mat.NewDense(2, 2, []float64{float64(button1.moveX), float64(button2.moveX), float64(button1.moveY), float64(button2.moveY)})
	detA := mat.Det(matrA)
	matrA1 := mat.NewDense(2, 2, []float64{float64(xTarget), float64(button2.moveX), float64(yTarget), float64(button2.moveY)})
	detA1 := mat.Det(matrA1)
	matrA2 := mat.NewDense(2, 2, []float64{float64(button1.moveX), float64(xTarget), float64(button1.moveY), float64(yTarget)})
	detA2 := mat.Det(matrA2)
	timesButton1 := detA1 / detA
	timesButton2 := detA2 / detA
	if !isFlat(timesButton1) || !isFlat(timesButton2) {
		return 0
	}
	tb1 := int(math.Round(timesButton1))
	tb2 := int(math.Round(timesButton2))
	if tb1 <= 100 && tb2 <= 100 && tb1 >= 0 && tb2 >= 0 {
		return button1.costs*tb1 + button2.costs*tb2
	}
	return 0
}

func calculateValues2(button1 Button, button2 Button, xTarget int, yTarget int) int {
	matrA := mat.NewDense(2, 2, []float64{float64(button1.moveX), float64(button2.moveX), float64(button1.moveY), float64(button2.moveY)})
	detA := mat.Det(matrA)
	matrA1 := mat.NewDense(2, 2, []float64{float64(xTarget), float64(button2.moveX), float64(yTarget), float64(button2.moveY)})
	detA1 := mat.Det(matrA1)
	matrA2 := mat.NewDense(2, 2, []float64{float64(button1.moveX), float64(xTarget), float64(button1.moveY), float64(yTarget)})
	detA2 := mat.Det(matrA2)
	timesButton1 := detA1 / detA
	timesButton2 := detA2 / detA
	if !isFlat(timesButton1) || !isFlat(timesButton2) {
		return 0
	}
	tb1 := int(math.Round(timesButton1))
	tb2 := int(math.Round(timesButton2))
	if tb1 >= 0 && tb2 >= 0 {
		return button1.costs*tb1 + button2.costs*tb2
	}
	return 0
}

func isFlat(num float64) bool {
	diff := math.Abs(math.Round(num) - num)
	return diff < 0.01
}

func createButton(line string) Button {
	line = strings.ReplaceAll(line, "X", "")
	line = strings.ReplaceAll(line, "Y", "")
	line = strings.ReplaceAll(line, "+", "")
	line = strings.ReplaceAll(line, ",", "")
	lineArr := strings.Split(line, " ")
	moveX, _ := strconv.Atoi(lineArr[2])
	moveY, _ := strconv.Atoi(lineArr[3])
	return Button{moveX: moveX, moveY: moveY}
}

func createTarget(line string) (int, int) {
	line = strings.ReplaceAll(line, "X", "")
	line = strings.ReplaceAll(line, "Y", "")
	line = strings.ReplaceAll(line, "=", "")
	line = strings.ReplaceAll(line, ",", "")
	lineArr := strings.Split(line, " ")
	x, _ := strconv.Atoi(lineArr[1])
	y, _ := strconv.Atoi(lineArr[2])
	return x, y
}
