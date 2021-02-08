package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 27. 移除元素
 */
func Test_leetcode_027(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	defer timeCost()()
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] != val {
			left++
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		right--
	}
	return left
}
