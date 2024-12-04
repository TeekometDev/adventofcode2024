package matrixhelpers

func CreateVertical(input [][]rune) [][]rune {
	rows := len(input)
	cols := len(input[0])
	rotated := make([][]rune, cols)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][i] = input[i][j]
		}
	}

	return rotated
}

func RotateMatrix45(matrix [][]rune) [][]rune {
	rows := len(matrix)
	cols := len(matrix[0])
	newRows := rows + cols - 1
	rotated := make([][]rune, newRows)
	for i := range rotated {
		rotated[i] = make([]rune, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[i+j][i] = matrix[i][j]
		}
	}
	return rotated
}

func RotateMatrix135(matrix [][]rune) [][]rune {
	rows := len(matrix)
	cols := len(matrix[0])
	rotated := make([][]rune, cols)
	for i := 0; i < rows; i++ {
		rotated[i] = make([]rune, rows)
		for j := 0; j < cols; j++ {
			rotated[i][j] = matrix[rows-1-j][i]
		}
	}
	return RotateMatrix45(rotated)
}
