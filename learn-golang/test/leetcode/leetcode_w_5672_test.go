package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5672. 检查数组是否经排序和轮转得到
 */

func Test_leetcode_5672(t *testing.T) {
	fmt.Println(check5672([]int{3, 4, 5, 1, 2}))
}

func check5672(nums []int) bool {
	defer timeCost()()
	length := len(nums)
	if length <= 1 {
		return true
	}
	// i: 轮转偏移
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			idx, idx2 := (i+j)%length, (i+j+1)%length
			if nums[idx] > nums[idx2] {
				break
			}
			if j == length-2 {
				return true
			}
		}
	}
	return false
}
