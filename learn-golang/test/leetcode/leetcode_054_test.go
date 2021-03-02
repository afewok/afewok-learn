package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 54. 螺旋矩阵
 */

func Test_leetcode_054(t *testing.T) {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	defer timeCost()()
	h1, w1, h2, w2, modle := 0, 0, len(matrix)-1, len(matrix[0])-1, 0
	result := make([]int, 0)
	for h1 <= h2 && w1 <= w2 {
		switch modle {
		case 0:
			for i := w1; i <= w2; i++ {
				result = append(result, matrix[h1][i])
			}
			h1++
		case 1:
			for i := h1; i <= h2; i++ {
				result = append(result, matrix[i][w2])
			}
			w2--
		case 2:
			for i := w2; i >= w1; i-- {
				result = append(result, matrix[h2][i])
			}
			h2--
		case 3:
			for i := h2; i >= h1; i-- {
				result = append(result, matrix[i][w1])
			}
			w1++
		}
		modle++
		if modle > 3 {
			modle = 0
		}
	}
	return result
}
