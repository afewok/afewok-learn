package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 136. 只出现一次的数字
 */
func Test_leetcode_136(t *testing.T) {
	fmt.Println([]int{2, 2, 1})
	fmt.Println([]int{4, 1, 2, 1, 2})
}

func singleNumber1(nums []int) int {
	defer timeCost()()
	mp := make(map[int]bool)
	for _, num := range nums {
		if _, OK := mp[num]; OK {
			delete(mp, num)
		} else {
			mp[num] = true
		}
	}
	for key := range mp {
		return key
	}
	return 0
}

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}
