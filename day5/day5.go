package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"teekometDev/filereader5"
)

func main() {
	result1 := Task1("input.txt")
	result2 := Task2("input.txt")
	fmt.Printf("RESULT T1: %d, RESULT T2: %d\n", result1, result2)
}

func Task1(fileName string) int {
	rawLines := filereader5.ReadFile(fileName)
	rules, updates := getPartLists(rawLines)
	filteredUpdates := filterValidLines(rules, updates)
	return calcResult1(filteredUpdates)
}

func Task2(fileName string) int {
	rawLines := filereader5.ReadFile(fileName)
	rules, updates := getPartLists(rawLines)
	invalids := invalidUpdates(&rules, &updates)
	var ordered [][]int
	for _, invalidRule := range invalids {
		ordered = append(ordered, orderLine(&rules, invalidRule))
	}
	return calcResult1(ordered)
}

func getPartLists(input []string) ([][]int, [][]int) {
	var rules [][]int
	var updates [][]int
	r_r, _ := regexp.Compile(`[0-9]+\|[0-9]+`)
	r_u, _ := regexp.Compile(`[0-9]+(,[0-9]+)+`)
	for _, line := range input {
		if r_r.FindString(line) != "" {
			command := strings.Split(line, "|")
			c_1, _ := strconv.Atoi(command[0])
			c_2, _ := strconv.Atoi(command[1])
			rules = append(rules, []int{c_1, c_2})
		} else if r_u.FindString(line) != "" {
			elements := strings.Split(line, ",")
			var temp_update []int
			for _, num := range elements {
				num_i, _ := strconv.Atoi(num)
				temp_update = append(temp_update, num_i)
			}
			updates = append(updates, temp_update)
		}
	}
	return rules, updates
}

func filterValidLines(ruleset [][]int, updates [][]int) [][]int {
	var validLines [][]int
	for _, update := range updates {
		isInvalid := false
		var forbiddenNumbers []int
		for _, num := range update {
			if isInvalid {
				continue
			}
			if slices.Contains(forbiddenNumbers, num) {
				isInvalid = true
				continue
			} else {
				forbiddenNumbers = append(forbiddenNumbers, returnRules(ruleset, num)...)
			}
		}
		if !isInvalid {
			validLines = append(validLines, update)
		}
	}
	return validLines
}

func invalidUpdates(ruleset *[][]int, updates *[][]int) [][]int {
	validLines := filterValidLines(*ruleset, *updates)
	var filtered [][]int
	for _, update := range *updates {
		found := false
		for _, validLine := range validLines {
			if found {
				continue
			} else if equalSlices(update, validLine) {
				found = true
			}
		}
		if !found {
			filtered = append(filtered, update)
		}
	}
	return filtered
}

func returnRules(ruleset [][]int, num int) []int {
	var invalidNumbers []int
	for _, rule := range ruleset {
		if rule[1] == num {
			invalidNumbers = append(invalidNumbers, rule[0])
		}
	}
	return invalidNumbers
}

func orderLine(rules *[][]int, line []int) []int {
	var forbiddenNumbers []int
	for i, num := range line {
		if slices.Contains(forbiddenNumbers, num) {
			line[i-1], line[i] = line[i], line[i-1]
			return orderLine(rules, line)
		} else {
			forbiddenNumbers = append(forbiddenNumbers, returnRules(*rules, num)...)
		}
	}
	return line
}

func calcResult1(validPages [][]int) int {
	result := 0
	for _, element := range validPages {
		middleIndex := int(len(element) / 2)
		result += element[middleIndex]
	}
	return result
}

func equalSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
