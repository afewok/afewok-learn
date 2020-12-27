package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 205. 同构字符串
 */

func Test_leetcode_205(t *testing.T) {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("aba", "baa"))
}

func isIsomorphic(s string, t string) bool {
	defer timeCost()()
	var list [2][255]byte
	var sByte, tByte byte
	length := len(s)
	for i := 0; i < length; i++ {
		tByte, sByte = list[0][s[i]], list[1][t[i]]
		if tByte == 0 && sByte == 0 {
			list[0][s[i]] = t[i]
			list[1][t[i]] = s[i]
		} else if tByte != t[i] || sByte != s[i] {
			return false
		}
	}
	return true
}

func isIsomorphic1(s string, t string) bool {
	defer timeCost()()
	mp1, mp2, length := make(map[byte]byte, 26), make(map[byte]byte, 26), len(s)
	for i := 0; i < length; i++ {
		t1, OK1 := mp1[s[i]]
		s1, OK2 := mp2[t[i]]
		if OK1 != OK2 || (OK1 && (t1 != t[i] || s[i] != s1)) {
			return false
		} else if !OK1 {
			mp1[s[i]] = t[i]
			mp2[t[i]] = s[i]
		}
	}
	return true
}
