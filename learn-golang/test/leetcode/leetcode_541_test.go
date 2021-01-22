package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 541. 反转字符串 II
 */

func Test_leetcode_541(t *testing.T) {
	fmt.Println(reverseStr("abcdefg", 2))
}

func reverseStr(s string, k int) string {
	defer timeCost()()
	length, result := len(s), ""
	direction, count := false, (length+k-1)/k
	for i := 0; i < count; i++ {
		s1, s2 := i*k, (i+1)*k
		if s2 > length {
			s2 = length
		}
		if direction {
			result += s[s1:s2]
		} else {
			result += reverseString541(s[s1:s2])
		}
		direction = !direction
	}
	return result
}

func reverseString541(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
