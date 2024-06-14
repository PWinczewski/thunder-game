package main

func CheckMatrixPattern(matrix [][]int64, i, j int, pattern [3][3]int64) bool {
	for di := 0; di < 3; di++ {
		for dj := 0; dj < 3; dj++ {
			if pattern[di][dj] != -1 && matrix[i+di][j+dj] != -1 && matrix[i+di][j+dj] != pattern[di][dj] {
				return false
			}
		}
	}
	return true
}

func PadMatrix(matrix [][]int64, borderValue int64) [][]int64 {
	rows := len(matrix)
	cols := len(matrix[0])

	paddedMatrix := make([][]int64, rows+2)
	for i := range paddedMatrix {
		paddedMatrix[i] = make([]int64, cols+2)
	}

	fillRow := make([]int64, cols+2)
	for j := range fillRow {
		fillRow[j] = borderValue
	}
	paddedMatrix[0] = fillRow
	paddedMatrix[rows+1] = fillRow

	for i := 1; i <= rows; i++ {
		paddedMatrix[i][0] = borderValue
		copy(paddedMatrix[i][1:cols+1], matrix[i-1])
		paddedMatrix[i][cols+1] = borderValue
	}

	return paddedMatrix
}

func MorphMatrix(matrix [][]int64, patterns [][3][3]int64, morphedValue int64, sweeps int) [][]int64 {
	rows := len(matrix)
	cols := len(matrix[0])

	paddedMatrix := PadMatrix(matrix, -1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, pattern := range patterns {
				if CheckMatrixPattern(paddedMatrix, i, j, pattern) {
					paddedMatrix[i+1][j+1] = morphedValue
					break
				}
			}
		}
	}

	morphedMatrix := make([][]int64, rows)
	for i := 1; i <= rows; i++ {
		morphedMatrix[i-1] = paddedMatrix[i][1 : cols+1]
	}

	if sweeps <= 1 {
		return morphedMatrix
	}

	return MorphMatrix(morphedMatrix, patterns, morphedValue, sweeps-1)
}
