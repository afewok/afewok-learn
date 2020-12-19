package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 88. 合并两个有序数组
 */
func Test_leetcode_088(t *testing.T) {
	arr1, arr2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	merge(arr1, 3, arr2, len(arr2))
	fmt.Println(arr1)

	arr1, arr2 = []int{0}, []int{1}
	merge(arr1, 0, arr2, 1)
	fmt.Println(arr1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	defer timeCost()()
	length, m, n := len(nums1)-1, m-1, n-1
	for n > -1 {
		if m > -1 && nums1[m] > nums2[n] {
			nums1[length] = nums1[m]
			m--
		} else {
			nums1[length] = nums2[n]
			n--
		}
		length--
	}
}
