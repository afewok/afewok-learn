package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 283. 移动零
 */

func Test_leetcode_5622(t *testing.T) {
	array := []int{0, 1, 0, 3, 12}
	moveZeroes(array)
	fmt.Println(array)
}

func moveZeroes(nums []int) {
	defer timeCost()()
	space, length := 0, len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			space++
		} else {
			nums[i-space], nums[i] = nums[i], nums[i-space]
		}
	}
}
