package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 59. 螺旋矩阵 II
 */

func Test_leetcode_059(t *testing.T) {
	fmt.Println(generateMatrix(4))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(1))
}

func generateMatrix(n int) [][]int {
	defer timeCost()()
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	x1, y1, x2, y2, sign, count := 0, 0, n-1, n-1, 0, 0
	for x1 <= x2 && y1 <= y2 {
		switch sign {
		case 0:
			for i := x1; i <= x2; i++ {
				count++
				result[y1][i] = count
			}
			y1++
		case 1:
			for i := y1; i <= y2; i++ {
				count++
				result[i][x2] = count
			}
			x2--
		case 2:
			for i := x2; i >= x1; i-- {
				count++
				result[y2][i] = count
			}
			y2--
		case 3:
			for i := y2; i >= y1; i-- {
				count++
				result[i][x1] = count
			}
			x1++
		}
		sign++
		if sign > 3 {
			sign = 0
		}
	}
	return result
}
