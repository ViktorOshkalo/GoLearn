package matrix

import "fmt"

type Matrix [][]int

type DiagonalType int

const (
	LeftToRight DiagonalType = 0
	RightToLeft DiagonalType = 1
)

func InitMatrix(rows int, columns int) Matrix {
	var matrix Matrix = make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, columns)
	}
	return matrix
}

func (matrix Matrix) GetRow(i int) []int {
	return GetRow(i, matrix)
}

func GetRow(i int, matrix Matrix) []int {
	return matrix[i]
}

func (matrix Matrix) GetColumn(j int) []int {
	return GetColumn(j, matrix)
}

func GetColumn(j int, matrix Matrix) []int {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][j]
	}
	return output
}

func (matrix Matrix) GetDiagonalLeftToRight() []int {
	return GetDiagonalLeftToRight(matrix)
}

func GetDiagonalLeftToRight(matrix Matrix) []int {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][i]
	}
	return output
}

func (matrix Matrix) GetDiagonalRightToLeft() []int {
	return GetDiagonalRightToLeft(matrix)
}

func GetDiagonalRightToLeft(matrix Matrix) []int {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][len(matrix)-1-i]
	}
	return output
}

func GetDiagonal(diagonalType int, matrix Matrix) []int {
	switch DiagonalType(diagonalType) {
	case LeftToRight:
		return GetDiagonalLeftToRight(matrix)
	case RightToLeft:
		return GetDiagonalRightToLeft(matrix)
	default:
		panic(fmt.Errorf("unknown diagonal type: %d", diagonalType))
	}
}

func GetSum(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}

func IsLineTaken(line []int) bool {
	return GetSum(line) > 0
}
