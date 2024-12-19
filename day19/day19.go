package main

import (
	"fmt"
	"strings"
	"teekometDev/filereader5"
)

func main() {
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file)
	res2 := Task2(file)
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res1, res2)
}

func Task1(input []string) int {
	patterns := createPatterns(input[0])
	var towels []string
	for i := 2; i < len(input); i++ {
		towels = append(towels, input[i])
	}
	res := 0
	dict := make(map[string]bool)
	for _, towel := range towels {
		if isPossibleTowel(towel, &patterns, &dict) {
			res++
		}
	}
	return res
}

func Task2(input []string) int {
	patterns := createPatterns(input[0])
	var towels []string
	for i := 2; i < len(input); i++ {
		towels = append(towels, input[i])
	}
	var solvableTowels []string
	dict := make(map[string]bool)
	for _, towel := range towels {
		if isPossibleTowel(towel, &patterns, &dict) {
			solvableTowels = append(solvableTowels, towel)
		}
	}
	res := 0
	solCache := make(map[string]int)
	for _, towel := range solvableTowels {
		val := trackPosibilities(towel, &patterns, &solCache)
		res += val
	}
	return res
}

func createPatterns(line string) []string {
	line = strings.ReplaceAll(line, " ", "")
	return strings.Split(line, ",")
}

func isPossibleTowel(towel string, patterns *[]string, dict *map[string]bool) bool {
	if len(towel) == 0 {
		return true
	}
	value, exists := (*dict)[towel]
	if exists {
		return value
	}
	for _, pattern := range *patterns {
		if strings.HasPrefix(towel, pattern) {
			tempRes := isPossibleTowel(towel[len(pattern):], patterns, dict)
			if tempRes {
				(*dict)[towel] = true
				return true
			}
		}
	}
	(*dict)[towel] = false
	return false
}

func trackPosibilities(towel string, patterns *[]string, cache *map[string]int) int {
	if len(towel) == 0 {
		return 1
	}
	value, exists := (*cache)[towel]
	if exists {
		return value
	}
	res := 0
	for _, pattern := range *patterns {
		if strings.HasPrefix(towel, pattern) {
			res += trackPosibilities(towel[len(pattern):], patterns, cache)
		}
	}
	(*cache)[towel] = res
	return res
}
