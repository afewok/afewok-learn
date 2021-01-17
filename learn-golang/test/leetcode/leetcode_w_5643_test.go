package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5243. 同积元组
 */

func Test_leetcode_5243(t *testing.T) {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))
	fmt.Println(tupleSameProduct([]int{1, 2, 4, 5, 10}))
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6, 8, 12}))
	fmt.Println(tupleSameProduct([]int{2, 3, 5, 7}))
}

func tupleSameProduct(nums []int) int {
	defer timeCost()()
	length, count, maps := len(nums), 0, make(map[int]int)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			maps[nums[i]*nums[j]]++
		}
	}
	for _, v := range maps {
		count += v * (v - 1) / 2
	}
	return count * 8
}
