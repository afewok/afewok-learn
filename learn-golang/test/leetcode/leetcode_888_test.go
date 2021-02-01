package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 888. 公平的糖果棒交换
 *
 */
func Test_leetcode_888(t *testing.T) {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))
	fmt.Println(fairCandySwap([]int{1, 2}, []int{2, 3}))
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))
}

func fairCandySwap(A []int, B []int) []int {
	defer timeCost()()
	totalA, totalB := 0, 0
	for _, v := range A {
		totalA += v
	}
	for _, v := range B {
		totalB += v
	}
	for _, a := range A {
		for _, b := range B {
			if totalA-a+b == totalB+a-b {
				return []int{a, b}
			}
		}
	}
	panic("panic")
}
