package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func Test_leetcode_290(t *testing.T) {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))

}

func wordPattern(pattern string, s string) bool {
	defer timeCost()()
	bytes, strs := []byte(pattern), strings.Split(s, " ")
	length1, length2 := len(bytes), len(strs)
	if length1 != length2 {
		return false
	}
	mp1, mp2 := make(map[byte]string), make(map[string]byte)
	for i := 0; i < length1; i++ {
		str, OK1 := mp1[bytes[i]]
		p, OK2 := mp2[strs[i]]
		if !OK1 && !OK2 {
			mp1[bytes[i]] = strs[i]
			mp2[strs[i]] = bytes[i]
		} else if !strings.EqualFold(strs[i], str) || pattern[i] != p {
			return false
		}
	}
	return true
}
