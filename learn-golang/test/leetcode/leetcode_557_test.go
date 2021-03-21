package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 557. 反转字符串中的单词 III
 */

func Test_leetcode_557(t *testing.T) {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	defer timeCost()()
	length := len(s)
	sub, result := 0, make([]byte, length)
	for i := 0; i < length; i++ {
		if s[i] != ' ' {
			continue
		}
		for j := i - 1; j >= sub; j-- {
			result[i-j-1+sub] = s[j]
		}
		result[i] = ' '
		sub = i + 1
	}
	for j := length - 1; j >= sub; j-- {
		result[length-j-1+sub] = s[j]
	}
	return string(result)
}
