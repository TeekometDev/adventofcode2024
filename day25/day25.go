package main

import (
	"fmt"
	"strings"
	"teekometDev/filereader5"
)

func main() {
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file)
	fmt.Printf("RESULT 1: %d\n", res1)
}

func Task1(file []string) int {
	keys := [][]int{}
	locks := [][]int{}
	currObj := []string{}
	for _, line := range file {
		if line == "" {
			createLockKey(currObj, &keys, &locks)
			currObj = []string{}
		} else {
			currObj = append(currObj, line)
		}
	}
	createLockKey(currObj, &keys, &locks)
	foundPairs := 0
	for _, lock := range locks {
		foundPairs += findKey(lock, &keys)
	}
	return foundPairs
}

func createLockKey(lines []string, keys *[][]int, locks *[][]int) {
	currArr := []int{}
	for i := 0; i < len(lines[0]); i++ {
		colCount := 0
		for j := 1; j < len(lines)-1; j++ {
			if lines[j][i] == '#' {
				colCount++
			}
		}
		currArr = append(currArr, colCount)
	}
	if strings.ContainsRune(lines[0], '#') {
		// Key
		*keys = append(*keys, currArr)
	} else {
		// Lock
		*locks = append(*locks, currArr)
	}
}

func findKey(lock []int, keys *[][]int) int {
	possibilities := 0
	for _, key := range *keys {
		stillValid := true
		for j := 0; j < len(key); j++ {
			if key[j]+lock[j] > 5 {
				stillValid = false
			}
		}
		if stillValid {
			// *keys = slices.Delete(*keys, i, i+1)
			possibilities++
		}
	}
	return possibilities
}
