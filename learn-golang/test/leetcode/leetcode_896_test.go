package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 896. 单调数列
 *
 */
func Test_leetcode_029(t *testing.T) {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(isMonotonic([]int{1, 3, 2}))
	fmt.Println(isMonotonic([]int{1, 2, 4, 5}))
	fmt.Println(isMonotonic([]int{1, 1, 1}))
}

func isMonotonic(A []int) bool {
	defer timeCost()()
	length, OK1, OK2 := len(A), true, true
	for i := 1; i < length; i++ {
		if OK2 && A[i-1] > A[i] {
			OK2 = false
		} else if OK1 && A[i-1] < A[i] {
			OK1 = false
		}
	}
	if OK1 || OK2 {
		return true
	}
	return false
}
