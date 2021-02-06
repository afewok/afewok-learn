package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5659. 删除字符串两端相同字符后的最短长度
 */

func Test_leetcode_5659(t *testing.T) {
	fmt.Println(minimumLength("aaa"))
	fmt.Println(minimumLength("bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"))
	fmt.Println(minimumLength("ca"))
	fmt.Println(minimumLength("cabaabac"))
	fmt.Println(minimumLength("aabccabba"))
}

func minimumLength(s string) int {
	defer timeCost()()
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			if left > 0 && s[left-1] == s[right] {
				right--
			} else if right < len(s)-1 && s[left] == s[right+1] {
				left++
			} else {
				return right - left + 1
			}
		} else {
			left++
			right--
		}
	}
	if right == left {
		if left > 0 && s[left-1] == s[right] {
			right--
		} else if right < len(s)-1 && s[left] == s[right+1] {
			left++
		} else {
			return right - left + 1
		}
	}
	return right - left + 1
}
