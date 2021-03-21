package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 77. 组合
 */

func Test_leetcode_077(t *testing.T) {
	fmt.Println(climbStairs(1))
}

func combine(n int, k int) [][]int {
	defer timeCost()()
	result := make([][]int, 0)
	for i := k; i < n; i++ {
		for j := k; j < n; j++ {

		}
	}
	return result
}
