package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 680. 验证回文字符串 Ⅱ
 */

func Test_leetcode_680(t *testing.T) {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}

func validPalindrome(s string) bool {
	defer timeCost()()
	f := func(left, right int, s string) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			if f(left+1, right, s) || f(left, right-1, s) {
				return true
			}
			return false
		}
		left++
		right--
	}
	return true
}
