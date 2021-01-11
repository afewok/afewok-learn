package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 58. 最后一个单词的长度
 */

func Test_leetcode_058(t *testing.T) {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	defer timeCost()()
	start, fLen, length := false, 0, len(s)
	for i := length - 1; i >= 0; i-- {
		if !start && s[i] == ' ' {
			continue
		}
		start = true
		if s[i] == ' ' {
			return fLen
		}
		fLen++
	}
	return fLen
}
