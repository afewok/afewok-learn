package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 941. 有效的山脉数组
 */
func Test_leetcode_941(t *testing.T) {
	fmt.Println(validMountainArray([]int{6, 7, 7, 8, 6}))
	// fmt.Println(validMountainArray([]int{2, 1}))
	// fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5, 1}))
}

func validMountainArray(arr []int) bool {
	defer timeCost()()
	length := len(arr)
	if length < 3 {
		return false
	} else if arr[0] >= arr[1] || arr[length-2] <= arr[length-1] {
		return false
	}
	max, sign := arr[0], true
	for i := 1; i < length; i++ {
		if arr[i] == arr[i-1] {
			return false
		}

		if sign {
			if max > arr[i] {
				sign = false
			} else {
				max = arr[i]
			}
		} else {
			if max <= arr[i] || arr[i-1] < arr[i] {
				return false
			}
		}
	}
	return true
}
