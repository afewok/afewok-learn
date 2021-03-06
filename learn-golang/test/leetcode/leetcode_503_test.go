package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 503. 下一个更大元素 II
 */
func Test_leetcode_503(t *testing.T) {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

func nextGreaterElements(nums []int) []int {
	defer timeCost()()
	result := make([]int, len(nums))
	for i, num := range nums {
		sub := -1
		for m := i + 1; m < len(nums); m++ {
			if num < nums[m] {
				sub = m
				break
			}
		}
		if sub == -1 {
			for m := 0; m < i; m++ {
				if num < nums[m] {
					sub = m
					break
				}
			}
		}
		if sub == -1 {
			result[i] = -1
		} else {
			result[i] = nums[sub]
		}
	}
	return result
}
