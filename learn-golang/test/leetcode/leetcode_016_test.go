package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 16. 最接近的三数之和
 */

func Test_leetcode_016(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func threeSumClosest(nums []int, target int) int {
	defer timeCost()()
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp > target {
				right--
			} else if tmp < target {
				left++
			} else {
				return target
			}
			if distance(tmp, target) < distance(res, target) {
				res = tmp
			}
		}
	}
	return res
}

func distance(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
