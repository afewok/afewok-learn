package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 421. 数组中两个数的最大异或值
 */

func Test_leetcode_421(t *testing.T) {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
	fmt.Println(findMaximumXOR([]int{0}))
	fmt.Println(findMaximumXOR([]int{2, 4}))
	fmt.Println(findMaximumXOR([]int{8, 10, 2}))
	fmt.Println(findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}))
}

func findMaximumXOR(nums []int) int {
	defer timeCost()()
	max := 0
	for _, v1 := range nums {
		for _, v2 := range nums {
			temp := v1 ^ v2
			if max < temp {
				max = temp
			}
		}
	}
	return max
}
