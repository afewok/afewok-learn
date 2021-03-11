package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 125. 验证回文串
 */

func Test_leetcode_125(t *testing.T) {
	fmt.Println(isPalindrome125("A man, a plan, a canal: Panama"))
}

func isPalindrome125(s string) bool {
	defer timeCost()()
	f := func(b byte) byte {
		if b >= 'a' && b <= 'z' {
			return b - 32
		}
		return b
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] < '0' || (s[left] > '9' && s[left] < 'A') || (s[left] > 'Z' && s[left] < 'a') || s[left] > 'z' {
			left++
		} else if s[right] < '0' || (s[right] > '9' && s[right] < 'A') || (s[right] > 'Z' && s[right] < 'a') || s[right] > 'z' {
			right--
		} else if f(s[left]) != f(s[right]) {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
