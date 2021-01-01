package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 239. 滑动窗口最大值
 */

func Test_leetcode_239(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
	fmt.Println(maxSlidingWindow([]int{9, 11}, 2))
	fmt.Println(maxSlidingWindow([]int{4, -2}, 2))
}

func maxSlidingWindow(nums []int, k int) []int {
	defer timeCost()()
	length, max, list := len(nums)-k, 0, make([]int, 0)
	for i := 0; i <= length; i++ {
		max = nums[i]
		for j := i + k - 1; j > i; j-- {
			if max < nums[j] {
				max = nums[j]
			}
		}
		list = append(list, max)
	}

	return list
}
