package matrix

import "fmt"

type DiagonalType int

const (
	LeftToRight DiagonalType = 0
	RightToLeft DiagonalType = 1
)

type Line []int

type Matrix [][]int

func InitMatrix(rows int, columns int) Matrix {
	var matrix Matrix = make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, columns)
	}
	return matrix
}

func (matrix Matrix) GetRow(i int) Line {
	return GetRow(i, matrix)
}

func (matrix Matrix) GetColumn(j int) Line {
	return GetColumn(j, matrix)
}

func (matrix Matrix) GetDiagonalLeftToRight() Line {
	return GetDiagonalLeftToRight(matrix)
}

func (matrix Matrix) GetDiagonalRightToLeft() Line {
	return GetDiagonalRightToLeft(matrix)
}

func GetRow(i int, matrix Matrix) Line {
	return matrix[i]
}

func (matrix Matrix) GetAllLines() []Line {
	var output []Line
	for i := range matrix {
		output = append(output, GetColumn(i, matrix))
		output = append(output, GetRow(i, matrix))
	}
	output = append(output, GetDiagonalLeftToRight(matrix))
	output = append(output, GetDiagonalRightToLeft(matrix))
	return output
}

func GetColumn(j int, matrix Matrix) Line {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][j]
	}
	return output
}

func GetDiagonalLeftToRight(matrix Matrix) Line {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][i]
	}
	return output
}

func GetDiagonalRightToLeft(matrix Matrix) Line {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][len(matrix)-1-i]
	}
	return output
}

func GetDiagonal(diagonalType int, matrix Matrix) Line {
	switch DiagonalType(diagonalType) {
	case LeftToRight:
		return GetDiagonalLeftToRight(matrix)
	case RightToLeft:
		return GetDiagonalRightToLeft(matrix)
	default:
		panic(fmt.Errorf("unknown diagonal type: %d", diagonalType))
	}
}

func GetSum(line Line) int {
	var sum int
	for _, val := range line {
		sum += val
	}
	return sum
}

func IsLinePartiallyTaken(line Line) bool {
	return GetSum(line) > 0
}

func IsLineFullyTaken(line Line) bool {
	return GetSum(line) == len(line)
}
