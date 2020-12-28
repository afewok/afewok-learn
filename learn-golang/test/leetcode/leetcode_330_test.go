package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 330. 按要求补齐数组
 */

func Test_leetcode_330(t *testing.T) {
	fmt.Println(minPatches([]int{1, 3}, 6))
	fmt.Println(minPatches([]int{1, 5, 10}, 20))
}

func minPatches(nums []int, n int) int {
	defer timeCost()()
	patches := 0
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x *= 2
			patches++
		}
	}
	return patches
}
