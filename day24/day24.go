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
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file)
	fmt.Printf("RESULT 1: %d\n", res1)
}

func Task1(file []string) int {
	initPart := true
	valueMap := make(map[string]bool)
	rules := []Rule{}
	for _, line := range file {
		if line == "" {
			initPart = false
			continue
		}
		if initPart {
			addInitVal(line, &valueMap)
		} else {
			addRule(line, &rules)
		}
	}
	applyRules(&rules, &valueMap)
	zReg, _ := regexp.Compile("z[0-9]+")
	keyArr := []string{}
	for key := range valueMap {
		if zReg.MatchString(key) {
			keyArr = append(keyArr, key)
		}
	}
	slices.Sort(keyArr)
	valueStr := ""
	for _, key := range keyArr {
		if valueMap[key] {
			valueStr = "1" + valueStr
		} else {
			valueStr = "0" + valueStr
		}
	}

	decVal, _ := strconv.ParseInt(valueStr, 2, 64)
	return int(decVal)
}

func addInitVal(line string, mapOfValues *map[string]bool) {
	line = strings.ReplaceAll(line, ":", "")
	lineArr := strings.Split(line, " ")
	if lineArr[1] == "1" {
		(*mapOfValues)[lineArr[0]] = true
	} else {
		(*mapOfValues)[lineArr[0]] = false
	}
}

func addRule(line string, listOfRules *[]Rule) {
	lineArr := strings.Split(line, " ")
	newRule := Rule{in1: lineArr[0], in2: lineArr[2], out: lineArr[4]}
	switch lineArr[1] {
	case "XOR":
		newRule.op = XOR
	case "OR":
		newRule.op = OR
	case "AND":
		newRule.op = AND
	}
	(*listOfRules) = append((*listOfRules), newRule)
}

func applyRules(rules *[]Rule, mapOfValues *map[string]bool) {
	currIndex := 0
	for len(*rules) > 0 {
		currRule := (*rules)[currIndex]
		val1, exists1 := (*mapOfValues)[currRule.in1]
		val2, exists2 := (*mapOfValues)[currRule.in2]
		if exists1 && exists2 {
			switch currRule.op {
			case XOR:
				(*mapOfValues)[currRule.out] = (val1 && !val2) || (!val1 && val2)
			case AND:
				(*mapOfValues)[currRule.out] = val1 && val2
			case OR:
				(*mapOfValues)[currRule.out] = val1 || val2
			}
			(*rules) = slices.Delete((*rules), currIndex, currIndex+1)
		} else {
			currIndex++
		}
		if currIndex >= len(*rules) {
			currIndex = 0
		}
	}
}
