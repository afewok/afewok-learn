package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1207. 独一无二的出现次数
 */

func Test_leetcode_1207(t *testing.T) {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}

func uniqueOccurrences(arr []int) bool {
	defer timeCost()()
	mp1, mp2 := make(map[int]int), make(map[int]bool)
	for _, num := range arr {
		mp1[num]++
	}
	for _, v := range mp1 {
		if _, OK := mp2[v]; OK {
			return false
		}
		mp2[v] = true
	}
	return true
}
