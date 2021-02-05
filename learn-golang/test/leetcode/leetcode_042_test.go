package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 42. 接雨水
 */

func Test_leetcode_042(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

func trap(height []int) int {
	defer timeCost()()
	result := 0

	return result
}
