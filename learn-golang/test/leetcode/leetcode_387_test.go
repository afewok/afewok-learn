package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 387. 字符串中的第一个唯一字符
 */

func Test_leetcode_387(t *testing.T) {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}

func firstUniqChar1(s string) int {
	defer timeCost()()
	list, length := make([]int, 26), len(s)+1
	for i, b := range s {
		v := list[b-97]
		if v == 0 {
			list[b-97] = i + 1
		} else if v > 0 {
			list[b-97] = -1
		}
	}
	result := length
	for i := 0; i < 26; i++ {
		if list[i] > 0 && result > list[i] {
			result = list[i]
		}
	}
	if result == length {
		return -1
	}
	return result - 1
}

func firstUniqChar(s string) int {
	defer timeCost()()
	mp := map[byte]int{}
	for i, b := range s {
		v, OK := mp[byte(b-97)]
		if !OK {
			mp[byte(b-97)] = i
		} else if v >= 0 {
			mp[byte(b-97)] = -1
		}
	}

	result := len(s)
	for _, v := range mp {
		if v >= 0 && result > v {
			result = v
		}
	}
	if result == len(s) {
		return -1
	}
	return result
}
