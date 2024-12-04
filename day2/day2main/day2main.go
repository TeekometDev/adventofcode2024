package main

import (
	"fmt"
	"math"
	"teekometDev/filereader"
	"time"
)

const CRITICAL_MAX_INC = 3
const CRITICAL_MIN_INC = 1

func safeReport(report []int) (bool, int) {
	var dec bool = false
	for index, number := range report {
		if index == 0 {
			if number > report[1] {
				dec = true
			}
			continue
		}
		var diff = math.Abs(float64(number - report[index-1]))
		if diff > CRITICAL_MAX_INC || diff < CRITICAL_MIN_INC {
			return false, index
		}
		if dec == true && (number > report[index-1]) {
			return false, index
		}
		if dec == false && (number < report[index-1]) {
			return false, index
		}
	}
	return true, -1
}

func safeReport2(report []int) bool {
	iterationLength := len(report)
	safe, _ := safeReport(report)
	if safe {
		return true
	}
	for i := 0; i < iterationLength; i++ {
		var tempArr []int
		for j := 0; j < iterationLength; j++ {
			if j != i {
				tempArr = append(tempArr, report[j])
			}
		}
		safe, _ = safeReport(tempArr)
		if safe {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()
	message := filereader.HelloWorld()
	fmt.Println((message))
	resultMatrix := filereader.ReadFile("input.txt")
	var safeReports int = 0
	var safeReports2 int = 0
	for _, line := range resultMatrix {
		safe, _ := safeReport(line)
		if safe {
			safeReports++
		}
		if safeReport2(line) {
			safeReports2++
		}
	}
	fmt.Print("Task1: ", safeReports)
	fmt.Print("Task2: ", safeReports2)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime: %v\n", duration.Seconds())
}
