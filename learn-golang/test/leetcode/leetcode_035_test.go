package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 35. 搜索插入位置
 */

func Test_leetcode_035(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) int {
	defer timeCost()()
	for i, num := range nums {
		if num == target || target < num {
			return i
		}
	}
	return len(nums)
}

// func searchInsert(nums []int, target int) int {
// 	defer timeCost()()
// 	left, right := 0, len(nums)-1
// 	mid := left + (right-left)/2 + 1
// 	for left <= right {
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] > target {
// 			left = mid
// 			mid = left + (right-left)/2 + 1
// 		} else if nums[mid] < target {
// 			right = mid
// 			mid = left + (right-left)/2 + 1
// 		}
// 	}

// }
