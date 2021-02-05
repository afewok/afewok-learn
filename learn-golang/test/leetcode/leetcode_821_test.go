package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 821. 字符的最短距离
 */

func Test_leetcode_821(t *testing.T) {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}

func shortestToChar(S string, C byte) []int {
	defer timeCost()()
	length := len(S)
	sub, result := -1, make([]int, length)
	for i := 0; i < length; i++ {
		if S[i] == C {
			if sub == -1 {
				for j := 0; j < i; j++ {
					result[j] = i - j
				}
			} else {
				for j := sub + 1; j < i; j++ {
					if j-sub < i-j {
						result[j] = j - sub
					} else {
						result[j] = i - j
					}
				}
			}
			sub = i
		}
	}
	if sub < length-1 {
		for j := sub + 1; j < length; j++ {
			result[j] = j - sub
		}
	}
	return result
}
