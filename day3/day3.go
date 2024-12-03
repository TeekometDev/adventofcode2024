package main

import (
	"fmt"
	"regexp"
	"strconv"
	"teekometDev/filereader3"
)

func Hello() string {
	return "Hello World"
}

func SolveT1(fileName string) int {
	result := 0
	text := filereader3.ReadFile(fileName)
	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	operations := r.FindAllString(text, -1)
	for _, operation := range operations {
		result += multiply(operation)
	}
	return result
}

func SolveT2(fileName string) int {
	result := 0
	multiply_enabled := true
	text := filereader3.ReadFile(fileName)
	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	// r_multi, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	r_enable, _ := regexp.Compile(`do\(\)`)
	r_disable, _ := regexp.Compile(`don't\(\)`)
	operations := r.FindAllString(text, -1)
	for _, operation := range operations {
		if r_enable.FindString(operation) != "" {
			multiply_enabled = true
			continue
		}
		if r_disable.FindString(operation) != "" {
			multiply_enabled = false
			continue
		}
		if multiply_enabled {
			result += multiply(operation)
		}
	}
	return result
}

func multiply(operation string) int {
	number_r, _ := regexp.Compile(`[0-9]+`)
	targets := number_r.FindAllString(operation, 2)
	num1, _ := strconv.Atoi(targets[0])
	num2, _ := strconv.Atoi(targets[1])
	return num1 * num2
}

func main() {
	result1 := SolveT1("input.txt")
	result2 := SolveT2("input.txt")
	fmt.Printf("Task1 Solution: %d\nTask2 Solution: %d\n", result1, result2)
}
