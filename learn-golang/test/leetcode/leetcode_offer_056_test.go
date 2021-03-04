package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 剑指 Offer 56 - II. 数组中数字出现的次数 II
 */
func Test_leetcode_offer_056(t *testing.T) {
	fmt.Println(singleNumber056([]int{3, 4, 3, 3}))
	fmt.Println(singleNumber056([]int{9, 1, 7, 9, 7, 9, 7}))
}

func singleNumber05601(nums []int) int {
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

func singleNumber05602(nums []int) int {
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

func singleNumber056(nums []int) int {
	defer timeCost()()
	seenOnce, seenTwice := 0, 0
	for _, num := range nums {
		seenOnce = ^seenTwice & (seenOnce ^ num)
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}
	return seenOnce
}
