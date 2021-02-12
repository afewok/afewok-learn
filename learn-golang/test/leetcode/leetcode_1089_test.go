package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1089. 复写零
 */

func Test_leetcode_1089(t *testing.T) {
	// arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	// arr := []int{1, 5, 2, 0, 6, 8, 0, 6, 0}
	arr := []int{0, 0, 0, 0, 0, 0, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
	arr = []int{1, 2, 3}
	duplicateZeros(arr)
	fmt.Println(arr)
}

func duplicateZeros(arr []int) {
	defer timeCost()()
	m, zero, length := 0, 0, len(arr)
	for i := 0; i < length && m+zero < length; i++ {
		if arr[i] == 0 {
			zero++
		}
		m++
	}
	m = m - 1
	sub := length - 1
	if m+zero == length {
		arr[sub] = arr[m]
		sub--
		m--
	}
	for i := m; i >= 0; i-- {
		if arr[i] == 0 {
			arr[sub] = arr[i]
			sub--
		}
		arr[sub] = arr[i]
		sub--
	}
}
