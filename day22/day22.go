package main

import (
	"fmt"
	"strconv"
	"teekometDev/filereader5"
)

func main() {
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file)
	res2 := Task2(file)
	fmt.Printf("RESULT 1: %d, RESULT 2: %d\n", res1, res2)
}

func Task1(file []string) int {
	secMap := make(map[int]int)
	retVal := 0
	for _, line := range file {
		num, _ := strconv.Atoi(line)
		calcSecret(2000, &num, &secMap)
		retVal += num
	}
	return retVal
}

func Task2(file []string) int {
	secMap := make(map[int]int)
	seqNumbers := make(map[Sequence]int)
	for _, line := range file {
		num, _ := strconv.Atoi(line)
		currSeqMap := determineSequences(num, 2000, &secMap)
		for key, value := range currSeqMap {
			seqNumbers[key] += value
		}
	}
	maxVal := 0
	for _, entry := range seqNumbers {
		if entry > maxVal {
			maxVal = entry
		}
	}
	return maxVal
}

type Sequence struct {
	zero  int
	one   int
	two   int
	three int
}

func determineSequences(startNumber int, times int, cache *map[int]int) map[Sequence]int {
	localMap := make(map[Sequence]int)
	tempVals := []int{getLastDigit(startNumber)}
	// Calculate first 4 Values
	for i := 0; i < 4; i++ {
		calcSecret(1, &startNumber, cache)
		tempVals = append(tempVals, getLastDigit(startNumber))
		times--
	}
	localMap[createSequence(tempVals)] = tempVals[4]
	for times > 1 {
		calcSecret(1, &startNumber, cache)
		tempVals = tempVals[1:]
		tempVals = append(tempVals, getLastDigit(startNumber))
		currSeq := createSequence(tempVals)
		_, exists := localMap[currSeq]
		// Add only if its the first time this sequence occurs
		if !exists {
			localMap[currSeq] = tempVals[4]
		}
		times--
	}
	return localMap
}

func calcSecret(times int, secret *int, secMap *map[int]int) {
	if times <= 0 {
		return
	}
	startNum := *secret
	sol, exists := (*secMap)[startNum]
	if exists {
		(*secret) = sol
		calcSecret(times-1, secret, secMap)
	} else {
		step1(secret)
		step2(secret)
		step3(secret)
		(*secMap)[startNum] = *secret
		calcSecret(times-1, secret, secMap)
	}
}

func step1(secret *int) {
	mul := (*secret) * 64
	mix(secret, mul)
	prune(secret)
}

func step2(secret *int) {
	div := (*secret) / 32
	mix(secret, div)
	prune(secret)
}

func step3(secret *int) {
	mul := (*secret) * 2048
	mix(secret, mul)
	prune(secret)
}

func mix(secret *int, itemToMix int) {
	(*secret) = (*secret) ^ itemToMix
}

func prune(secret *int) {
	(*secret) = (*secret) % 16777216
}

func getLastDigit(num int) int {
	return num % 10
}

func createSequence(inputNums []int) Sequence {
	retSeq := Sequence{}
	retSeq.zero = inputNums[1] - inputNums[0]
	retSeq.one = inputNums[2] - inputNums[1]
	retSeq.two = inputNums[3] - inputNums[2]
	retSeq.three = inputNums[4] - inputNums[3]
	return retSeq
}
