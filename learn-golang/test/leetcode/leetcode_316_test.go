package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 316. 去除重复字母
 */

func Test_leetcode_316(t *testing.T) {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}

func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}
