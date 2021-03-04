package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 137. 只出现一次的数字 II
 */
func Test_leetcode_137(t *testing.T) {
	fmt.Println(singleNumber137([]int{3, 4, 3, 3}))
	fmt.Println(singleNumber137([]int{9, 1, 7, 9, 7, 9, 7}))
}

func singleNumber13701(nums []int) int {
	defer timeCost()()
	result, temp, mp := 0, 0, make(map[int]bool)
	for _, num := range nums {
		temp += num
		if _, OK := mp[num]; !OK {
			mp[num] = true
		}
	}
	for k := range mp {
		result += k
	}
	return (result*3 - temp) / 2
}

func singleNumber13702(nums []int) int {
	defer timeCost()()
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for k, v := range mp {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumber137(nums []int) int {
	seenOnce, seenTwice := 0, 0
	for _, num := range nums {
		seenOnce = ^seenTwice & (seenOnce ^ num)
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}
	return seenOnce
}
