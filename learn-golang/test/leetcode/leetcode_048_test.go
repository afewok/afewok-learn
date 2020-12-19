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
	rotate(data1)
	fmt.Println(data1)
	data2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}
	rotate(data2)
	fmt.Println(data2)
}

func rotate(matrix [][]int) {
	defer timeCost()()
	length := len(matrix)
	x, y, max, half, temp := 0, 0, length-1, length/2, 0
	for y < half {
		temp = matrix[y][x]
		matrix[y][x] = matrix[max-x][y]
		matrix[max-x][y] = matrix[max-y][max-x]
		matrix[max-y][max-x] = matrix[x][max-y]
		matrix[x][max-y] = temp
		if x < max-y-1 {
			x++
		} else {
			y++
			x = y
		}
	}
}
