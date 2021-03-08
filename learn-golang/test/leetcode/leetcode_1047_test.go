package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1047. 删除字符串中的所有相邻重复项
 */

func Test_leetcode_1047(t *testing.T) {
	fmt.Println(removeDuplicates1047("aaaaaaaaa"))
	fmt.Println(removeDuplicates1047("abbaca"))
}

func removeDuplicates1047(S string) string {
	defer timeCost()()
	length, result := len(S), make([]byte, 0)
	for i := 0; i < length; i++ {
		if len(result) > 0 && result[len(result)-1] == S[i] {
			result = result[:len(result)-1]
		} else {
			result = append(result, S[i])
		}
	}
	return string(result)
}
