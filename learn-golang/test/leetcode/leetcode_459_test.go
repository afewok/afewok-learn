package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 459. 重复的子字符串
 */

func Test_leetcode_459(t *testing.T) {
	fmt.Println(repeatedSubstringPattern("bb"))
	fmt.Println(repeatedSubstringPattern("abac"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}

func repeatedSubstringPattern(s string) bool {
	defer timeCost()()
	length, count, pattern := len(s), 2, false
	for length >= count {
		if length%count == 0 {
			step := length / count
			temp := s[0:step]
			pattern = true
			for i := 2; i <= count; i++ {
				if temp != s[step*(i-1):step*i] {
					pattern = false
					break
				}
			}
			if pattern {
				return pattern
			}
		}
		count++
	}
	return false
}
