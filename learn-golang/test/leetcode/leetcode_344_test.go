package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 344. 反转字符串
 */

func Test_leetcode_344(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte) {
	defer timeCost()()
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
