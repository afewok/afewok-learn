package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1243. 数组变换
 */

func Test_leetcode_1243(t *testing.T) {
	// fmt.Println(transformArray([]int{6, 2, 3, 4}))
	// fmt.Println(transformArray([]int{2, 1, 2, 1, 1, 2, 2, 1}))
	// fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5}))
	fmt.Println(transformArray([]int{6, 5, 8, 6, 7, 7, 3, 9, 8, 8, 3, 1, 2, 9, 8, 3}))
	// fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5}))
}

func transformArray(arr []int) []int {
	defer timeCost()()
	OK, temp, length := true, arr[0], len(arr)-1
	for {
		OK = true
		temp = arr[0]
		for i := 1; i < length; i++ {
			if arr[i] < temp && arr[i] < arr[i+1] {
				temp = arr[i]
				arr[i]++
				OK = false
			} else if arr[i] > temp && arr[i] > arr[i+1] {
				temp = arr[i]
				arr[i]--
				OK = false
			} else {
				temp = arr[i]
			}
		}

		if OK {
			break
		}
	}
	return arr
}
