package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1064. 不动点
 */

func Test_leetcode_1064(t *testing.T) {
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3, 7}))
	fmt.Println(fixedPoint([]int{0, 2, 5, 8, 17}))
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9}))
}

func fixedPoint1(arr []int) int {
	defer timeCost()()
	for i, v := range arr {
		if i == v {
			return i
		}
	}
	return -1
}

func fixedPoint(arr []int) int {
	defer timeCost()()
	left, mid, right := 0, 0, len(arr)-1
	for left <= right {
		mid = left + (right-left)/2
		if mid == arr[mid] {
			right = mid - 1
		} else if mid > arr[mid] {
			left = mid + 1
		} else if mid < arr[mid] {
			right = mid - 1
		}
	}
	if left == len(arr) || arr[left] != left {
		return -1
	}
	return left
}
