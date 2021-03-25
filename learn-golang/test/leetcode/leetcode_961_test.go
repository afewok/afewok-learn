package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 961. 重复 N 次的元素
 */
func Test_leetcode_961(t *testing.T) {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}

func repeatedNTimes(A []int) int {
	defer timeCost()()
	mp := make(map[int]bool)
	for _, item := range A {
		if _, OK := mp[item]; OK {
			return item
		}
		mp[item] = true
	}

	panic("")
}
