package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 剑指 Offer 61. 扑克牌中的顺子
 */
func Test_leetcode_offer_061(t *testing.T) {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
}

func isStraight(nums []int) bool {
	defer timeCost()()
	mp, max, min := make(map[int]bool, 5), 0, 14
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			continue
		} else if _, OK := mp[nums[i]]; OK {
			return false
		} else {
			mp[nums[i]] = true
			if max < nums[i] {
				max = nums[i]
			}
			if min > nums[i] {
				min = nums[i]
			}
		}
	}
	return max-min < 5
}
