package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 859. 亲密字符串
 */

func Test_leetcode_859(t *testing.T) {
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("ab", "ab"))
	fmt.Println(buddyStrings("aa", "aa"))
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
	fmt.Println(buddyStrings("", "aa"))
}

func buddyStrings(A string, B string) bool {
	defer timeCost()()
	lenA, lenB, mp, sub1, sub2, result := len(A), len(B), make(map[byte]int), -1, -1, false
	if lenA != lenB || lenA == 0 {
		return false
	}

	for i := 0; i < lenA; i++ {
		mp[A[i]]++
		if A[i] == B[i] {
			continue
		} else if sub1 == -1 {
			sub1 = i
		} else if sub2 == -1 {
			sub2 = i
			if A[sub1] == B[sub2] && A[sub2] == B[sub1] {
				result = true
			}
		} else {
			return false
		}
	}

	if result {
		return result
	} else if sub1 != -1 {
		return false
	}
	for _, v := range mp {
		if v > 1 {
			return true
		}
	}
	return false
}
