package filereader5

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ReadFile(fileName string) []string {
	start := time.Now()
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	var returnArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		returnArray = append(returnArray, text)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("File Runtime: %v\n", duration.Seconds())
	return returnArray
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
