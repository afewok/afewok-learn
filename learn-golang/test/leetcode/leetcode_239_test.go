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
	if len(nums) == 0 || k == 0 || k == 1 {
		return nums
	}
	queue := make([]int, 0, k)
	result := make([]int, len(nums)-k+1)
	for i, v := range nums {
		for len(queue) != 0 && v >= nums[queue[len(queue)-1]] {
			queue = queue[0 : len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k && queue[0] == i-k {
			queue = queue[1:]
		}
		if i >= k-1 {
			result[i-k+1] = nums[queue[0]]
		}
	}
	return result
}
