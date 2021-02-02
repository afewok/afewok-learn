package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 832. 翻转图像
 */

func Test_leetcode_832(t *testing.T) {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}

func flipAndInvertImage(A [][]int) [][]int {
	defer timeCost()()
	f := func(a int) int {
		if a == 0 {
			return 1
		}
		return 0
	}
	hLen, wLen := len(A), len(A[0])
	half := wLen / 2
	for i := 0; i < hLen; i++ {
		for j := 0; j < half; j++ {
			A[i][j], A[i][wLen-j-1] = f(A[i][wLen-j-1]), f(A[i][j])
		}
	}
	if half*2 < wLen {
		for i := 0; i < hLen; i++ {
			A[i][half] = f(A[i][half])
		}
	}
	return A
}
