package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 867. 转置矩阵
 *
 */
func Test_leetcode_867(t *testing.T) {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}))
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6}}))
}

func transpose(A [][]int) [][]int {
	defer timeCost()()
	w, h := len(A), len(A[0])
	B := make([][]int, h)
	for i := 0; i < h; i++ {
		B[i] = make([]int, w)
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			B[j][i] = A[i][j]
		}
	}
	return B
}
