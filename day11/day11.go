package main

import (
	"fmt"
	"strconv"
	"strings"
	"teekometDev/filereader3"
)

func main() {
	input := filereader3.ReadFile("input.txt")
	res_1 := Task1(input)
	res_2 := Task2(input)
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res_1, res_2)
}

func Task1(line string) int {
	str_arr := strings.Split(line, " ")
	for i := 0; i < 25; i++ {
		var newArr []string
		for _, element := range str_arr {
			applyRules(element, &newArr)
		}
		str_arr = newArr
	}
	return len(str_arr)
}

func Task2(line string) int {
	str_arr := strings.Split(line, " ")
	val := 0
	stoneMap := make(map[int]int)
	for _, stone := range str_arr {
		value, _ := strconv.Atoi(stone)
		stoneMap[value] = stoneMap[value] + 1
	}
	for i := 0; i < 75; i++ {
		stoneMap = calcNewStonemap(stoneMap)
	}
	for _, count := range stoneMap {
		val += count
	}
	return val
}

func applyRules(actNum string, arr *[]string) {
	numAsInt, _ := strconv.Atoi(actNum)
	if numAsInt == 0 {
		(*arr) = append((*arr), "1")
	} else if len(actNum)%2 == 0 {
		actLen := len(actNum)
		num1_s := actNum[:actLen/2]
		num2_s := actNum[actLen/2:]
		num1_i, _ := strconv.Atoi(num1_s)
		num2_i, _ := strconv.Atoi(num2_s)
		num1_s = strconv.Itoa(num1_i)
		num2_s = strconv.Itoa(num2_i)
		(*arr) = append((*arr), num1_s)
		(*arr) = append((*arr), num2_s)
	} else {
		num, _ := strconv.Atoi(actNum)
		num *= 2024
		num_s := strconv.Itoa(num)
		(*arr) = append((*arr), num_s)
	}
}

func calcNewStonemap(stoneMap map[int]int) map[int]int {
	nextMap := make(map[int]int)
	for stone, count := range stoneMap {
		var newStones []int
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if digits := strconv.Itoa(stone); len(digits)%2 == 0 {
			a, _ := strconv.Atoi(digits[:len(digits)/2])
			b, _ := strconv.Atoi(digits[len(digits)/2:])
			newStones = append(newStones, a)
			newStones = append(newStones, b)
		} else {
			newStones = append(newStones, stone*2024)
		}
		for _, newStone := range newStones {
			nextMap[newStone] = nextMap[newStone] + count
		}
	}
	return nextMap
}
