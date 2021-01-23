package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 704. 二分查找
 *
 */
func Test_leetcode_704(t *testing.T) {
	fmt.Println(search2([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func search2(nums []int, target int) int {
	defer timeCost()()
	length := len(nums)
	if length < 1 {
		return -1
	}
	left, right := 0, length-1
	for left <= right {
		temp := (right-left)/2 + left
		if nums[temp] == target {
			return temp
		} else if nums[temp] < target {
			left = temp + 1
		} else if nums[temp] > target {
			right = temp - 1
		}
	}
	return -1
}
