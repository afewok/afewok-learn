package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5657. 唯一元素的和
 */

func Test_leetcode_5657(t *testing.T) {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 5}))
}

func sumOfUnique(nums []int) int {
	defer timeCost()()
	max, mp := 0, make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for k, v := range mp {
		if v == 1 {
			max += k
		}
	}
	return max
}
