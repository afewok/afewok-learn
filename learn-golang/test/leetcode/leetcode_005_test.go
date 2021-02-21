package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5. 最长回文子串
 */
func Test_leetcode_005(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
}

func longestPalindrome(s string) string {
	defer timeCost()()
	return ""
}
