package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 766. 托普利茨矩阵
 *
 * 思路：
 */
func Test_leetcode_766(t *testing.T) {
	fmt.Println(isToeplitzMatrix([][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2}}))

	fmt.Println(isToeplitzMatrix([][]int{
		{1, 2, 3},
		{5, 1, 2},
		{9, 5, 1},
		{7, 9, 5}}))

	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}

func isToeplitzMatrix(matrix [][]int) bool {
	defer timeCost()()
	heightMax, widthMax, num := len(matrix), len(matrix[0]), 0
	for i := len(matrix) - 2; i > 0; i-- {
		num = matrix[i][0]
		for j := 1; j+i < heightMax && j < widthMax; j++ {
			if num != matrix[j+i][j] {
				return false
			}
		}
	}
	for i := 0; i < widthMax; i++ {
		num = matrix[0][i]
		for j := 1; j < heightMax && j+i < widthMax; j++ {
			if num != matrix[j][j+i] {
				return false
			}
		}
	}
	return true
}
