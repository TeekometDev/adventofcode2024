package filereader5

import (
	"bufio"
	"os"
)

func ReadFile(fileName string) []string {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	var returnArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		returnArray = append(returnArray, text)
	}
	return returnArray
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
