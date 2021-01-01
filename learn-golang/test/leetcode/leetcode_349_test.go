package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 349. 两个数组的交集
 */

func Test_leetcode_349(t *testing.T) {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersection(nums1 []int, nums2 []int) []int {
	defer timeCost()()
	mp1, mp2, list := make(map[int]bool, len(nums1)), make(map[int]bool, len(nums2)), make([]int, 0)
	for _, v := range nums1 {
		mp1[v] = true
	}
	for _, v := range nums2 {
		mp2[v] = true
	}

	for k := range mp1 {
		if _, OK := mp2[k]; OK {
			list = append(list, k)
		}
	}
	return list
}
