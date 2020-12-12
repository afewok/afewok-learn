package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 217. 存在重复元素
 */
func Test_leetcode_217(t *testing.T) {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	fmt.Println(containsDuplicate([]int{1, 5, -2, -4, 0}))
	fmt.Println(containsDuplicate([]int{-1200000005, -1200000005}))
}

func containsDuplicate(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}
	mp := make(map[int]byte, length)
	for i := 0; i < length; i++ {
		if _, ok := mp[nums[i]]; ok {
			return true
		}
		mp[nums[i]] = 0
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}
	var temp, flag1, flag2 uint = 0, 0, 0
	for _, num := range nums {
		if num < 0 {
			temp = 1 << -num
			if temp&flag2 > 0 {
				return true
			}
			flag2 = flag2 | temp
		} else {
			temp = 1 << num
			if temp&flag1 > 0 {
				return true
			}
			flag1 = flag1 | temp
		}
	}
	return false
}
