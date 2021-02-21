package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5685. 交替合并字符串
 */

func Test_leetcode_5685(t *testing.T) {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	defer timeCost()()
	left, right, len1, len2, result := 0, 0, len(word1), len(word2), make([]byte, 0)
	for left < len1 || right < len2 {
		if left < len1 && right < len2 {
			result = append(result, word1[left], word2[right])
			left++
			right++
		} else if left < len1 {
			result = append(result, word1[left:]...)
			break
		} else {
			result = append(result, word2[right:]...)
			break
		}
	}
	return string(result)
}
