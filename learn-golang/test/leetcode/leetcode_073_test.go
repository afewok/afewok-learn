package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 73. 矩阵置零
 */

func Test_leetcode_073(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	defer timeCost()()
	xList, yList := make([]int, 0), make([]int, 0)
	for y, items := range matrix {
		for x, item := range items {
			if item == 0 {
				xList = append(xList, x)
				yList = append(yList, y)
			}
		}
	}

	for i := range yList {
		for m := 0; m < len(matrix); m++ {
			matrix[m][xList[i]] = 0
		}
		for n := 0; n < len(matrix[0]); n++ {
			matrix[yList[i]][n] = 0
		}
	}
}
