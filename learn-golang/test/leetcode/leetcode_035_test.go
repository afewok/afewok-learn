package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 35. 搜索插入位置
 */

func Test_leetcode_035(t *testing.T) {
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert1(nums []int, target int) int {
	defer timeCost()()
	for i, num := range nums {
		if num == target || target < num {
			return i
		}
	}
	return len(nums)
}

func searchInsert(nums []int, target int) int {
	defer timeCost()()
	left, right := 0, len(nums)-1
	if target < nums[0] {
		return 0
	} else if target > nums[right] {
		return right + 1
	}

	for left <= right {
		mid := left + (right-left+1)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return left
}
