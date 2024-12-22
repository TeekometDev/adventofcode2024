package main

import (
	"fmt"
	"strconv"
	"teekometDev/filereader5"
)

func main() {
	file := filereader5.ReadFile("input.txt")
	res1 := Task1(file)
	fmt.Printf("RESULT 1: %d\n", res1)
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
