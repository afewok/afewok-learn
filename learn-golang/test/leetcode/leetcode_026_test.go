package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 26. 删除排序数组中的重复项
 *
 * 思路：双指针法
 */
func Test_leetcode_026(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	sub := 0
	for _, num := range nums {
		if nums[sub] < num {
			sub++
			nums[sub] = num
		}
	}
	return sub + 1
}
