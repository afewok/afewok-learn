package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 48. 旋转图像
 *
 * 思路：
 */
func Test_leetcode_048(t *testing.T) {
	data1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	data2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}
	rotate(data1)
	rotate(data2)
	fmt.Println(data1)
	fmt.Println(data2)
}

func rotate(matrix [][]int) {
	defer timeCost()()
	length := len(matrix)
	x1, y1, x2, y2, temp := 0, 0, length-1, length-1, 0
	for x1 < x2 && y1 < y2 {
		temp = matrix[y1][x1]
		matrix[y1][x1] = matrix[y2][x1]
		matrix[y2][x1] = matrix[y2][x2]
		matrix[y2][x2] = matrix[y1][x2]
		matrix[y1][x2] = temp
		if x1 < x2 {
			x1++
		} else {
			y1++
			x1 = y1
		}

		if y1 < y2 {
			y2--
		} else {
			x2--
			y2 = x2
		}
	}
}
