package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 304. 二维区域和检索 - 矩阵不可变
 */

func Test_leetcode_304(t *testing.T) {

	numMatrix := Constructor304([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(numMatrix)
	fmt.Println(numMatrix.SumRegion(2, 1, 4, 3))
	fmt.Println(numMatrix.SumRegion(1, 1, 2, 2))
	fmt.Println(numMatrix.SumRegion(1, 2, 2, 4))
}

type NumMatrix struct {
	matrix [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	lenH := len(matrix)
	if lenH == 0 {
		return NumMatrix{make([][]int, 0)}
	}
	lenW := len(matrix[0])
	for i := 0; i < lenH; i++ {
		for j := 0; j < lenW; j++ {
			temp := 0
			if j > 0 {
				temp = matrix[i][j-1]
			}
			matrix[i][j] += temp
		}
	}
	return NumMatrix{matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row2 < 0 || col2 < 0 || len(this.matrix) <= 0 {
		return 0
	}

	if row1 < 0 {
		row1 = 0
	}
	if col1 < 0 {
		col1 = 0
	}
	if row2 >= len(this.matrix) {
		row2 = len(this.matrix) - 1
	}
	if col2 >= len(this.matrix[0]) {
		col2 = len(this.matrix[0]) - 1
	}
	result := 0
	for i := row1; i <= row2; i++ {
		result += this.matrix[i][col2]
		if col1 > 0 {
			result -= this.matrix[i][col1-1]
		}
	}
	return result
}
