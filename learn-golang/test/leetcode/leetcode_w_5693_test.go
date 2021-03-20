package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5693. 字符串中第二大的数字
 */

func Test_leetcode_5693(t *testing.T) {
	fmt.Println(secondHighest("dfa12321afd"))
	fmt.Println(secondHighest("abc1111"))
}

func secondHighest(s string) int {
	defer timeCost()()
	mp1 := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	mp2, arr := make(map[rune]bool), make([]int, 0)
	for _, item := range s {
		if v, OK1 := mp1[item]; OK1 {
			if _, OK2 := mp2[item]; !OK2 {
				mp2[item] = true
				arr = append(arr, v)
			}
		}
	}
	length := len(arr)
	if length < 2 {
		return -1
	}
	max1, max2 := arr[0], arr[1]
	if max1 < max2 {
		max1, max2 = max2, max1
	}
	for i := 2; i < length; i++ {
		if max1 < arr[i] {
			max2, max1 = max1, arr[i]
		} else if max2 < arr[i] {
			max2 = arr[i]
		}
	}
	return max2
}
