package filereader3

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(fileName string) string {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	return text
}
