package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1662. 检查两个字符串数组是否相等
 */

func Test_leetcode_1662(t *testing.T) {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
	fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddef"}))
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "de"}, []string{"abcddef"}))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	defer timeCost()()
	arr, sub := make([]rune, 0), 0
	for _, items := range word1 {
		for _, item := range items {
			arr = append(arr, item)
		}
	}
	length := len(arr)
	for _, items := range word2 {
		for _, item := range items {
			if sub >= length || arr[sub] != item {
				return false
			}
			sub++
		}
	}
	if len(arr) == sub {
		return true
	}
	return false
}
