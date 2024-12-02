package filereader

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func HelloWorld() string {
	return "Hello World!"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	var matrix [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var row []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			check(err)
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
