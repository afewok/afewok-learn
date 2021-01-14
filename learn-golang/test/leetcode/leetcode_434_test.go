package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 434. 字符串中的单词数
 */

func Test_leetcode_434(t *testing.T) {
	fmt.Println(countSegments("  Hello, my name is John  "))
}

func countSegments(s string) int {
	defer timeCost()()
	count, temp := 0, false
	for _, b := range s {
		if b != ' ' && !temp {
			temp = true
			count++
		} else if b == ' ' && temp {
			temp = false
		}
	}
	return count
}
