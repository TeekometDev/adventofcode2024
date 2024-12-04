package filereader4

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileAsMatrix(filename string) [][]rune {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	var matrix [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}
	check(scanner.Err())

	return matrix
}
