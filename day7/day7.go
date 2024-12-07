package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"teekometDev/filereader5"
)

func main() {
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file)
	res2 := Task2(file)
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res1, res2)
}

func Task1(file []string) int {
	result := 0
	for _, line := range file {
		ex_res, elList := generateElements(line)
		if checkCalculation(ex_res, elList) {
			result += ex_res
		}
	}
	return result
}

func Task2(file []string) int {
	result := 0
	for _, line := range file {
		ex_res, elList := generateElements(line)
		if calculateBase3(ex_res, elList) {
			result += ex_res
		}
	}
	return result
}

func generateElements(line string) (int, []int) {
	element_s := strings.Split(line, " ")
	exRes, _ := strconv.Atoi(strings.Replace(element_s[0], ":", "", -1))
	var elementList []int
	for i := range element_s {
		if i > 0 {
			el, _ := strconv.Atoi(element_s[i])
			elementList = append(elementList, el)
		}
	}
	return exRes, elementList
}

func checkCalculation(expectedResult int, numbers []int) bool {
	opLen := int(math.Pow(2, float64(len(numbers)-1)))
	for i := range opLen {
		binList := createOpList(i, len(numbers)-1)
		if calculate(numbers, binList) == expectedResult {
			return true
		}
	}
	return false
}

func calculateBase3(expectedResult int, numbers []int) bool {
	opLen := int(math.Pow(3, float64(len(numbers)-1)))
	for i := range opLen {
		tList := create3List(i, len(numbers)-1)
		if calc3(numbers, tList) == expectedResult {
			return true
		}
	}
	return false
}

func calculate(numbers []int, operationList []bool) int {
	res := numbers[0]
	for i := 0; i < len(operationList); i++ {
		if operationList[i] {
			res *= numbers[i+1]
		} else {
			res += numbers[i+1]
		}
	}
	return res
}

func calc3(numbers []int, operationList []byte) int {
	res := numbers[0]
	for i := 0; i < len(operationList); i++ {
		el := operationList[i]
		if el == 0 {
			res += numbers[i+1]
		} else if el == 1 {
			res *= numbers[i+1]
		} else {
			log := math.Log10(float64(numbers[i+1]))
			pow := int(math.Floor(log))
			res *= int(math.Pow10(pow + 1))
			res += numbers[i+1]
		}
	}
	return res
}

func createOpList(opNum int, length int) []bool {
	binStr := strconv.FormatInt(int64(opNum), 2)
	if len(binStr) < length {
		binStr = strings.Repeat("0", length-len(binStr)) + binStr
	}

	var boolArray []bool
	for _, char := range binStr {
		if char == '1' {
			boolArray = append(boolArray, true)
		} else {
			boolArray = append(boolArray, false)
		}
	}
	return boolArray
}

func create3List(opNum int, length int) []byte {
	tStr := strconv.FormatInt(int64(opNum), 3)
	if len(tStr) < length {
		tStr = strings.Repeat("0", length-len(tStr)) + tStr
	}

	var numArray []byte
	for _, char := range tStr {
		if char == '1' {
			numArray = append(numArray, byte(1))
		} else if char == '2' {
			numArray = append(numArray, byte(2))
		} else {
			numArray = append(numArray, byte(0))
		}
	}
	return numArray
}
