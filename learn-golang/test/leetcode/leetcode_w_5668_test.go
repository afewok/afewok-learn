package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5668. 最长的美好子字符串
 */

func Test_leetcode_5668(t *testing.T) {
	fmt.Println(longestNiceSubstring("wWOExoVhvXebB"))
	fmt.Println(longestNiceSubstring("abABB"))
	fmt.Println(longestNiceSubstring("YazaAay"))
	fmt.Println(longestNiceSubstring("Bb"))
	fmt.Println(longestNiceSubstring("c"))
	fmt.Println(longestNiceSubstring("dDzeE"))
}

func longestNiceSubstring(s string) string {
	defer timeCost()()
	results := make([]string, 0)
	getString5668(s, &results)
	result := results[0]
	for _, v := range results {
		if len(result) < len(v) {
			result = v
		}
	}
	return result
}

func getString5668(s string, results *[]string) {
	mp1, mp2, sub, strs := make(map[rune]bool), make(map[rune]bool), 0, make([]string, 0)
	for _, c := range s {
		mp1[c] = true
	}
	for k := range mp1 {
		if k < 97 {
			if _, OK := mp1[k+32]; !OK {
				mp2[k] = true
			}
		} else {
			if _, OK := mp1[k-32]; !OK {
				mp2[k] = true
			}
		}
	}

	for i, c := range s {
		if _, OK := mp2[c]; OK {
			temp := s[sub:i]
			if temp != "" {
				strs = append(strs, temp)
			}
			sub = i + 1
		}
	}
	temp := s[sub:]
	if temp != "" || len(strs) == 0 {
		strs = append(strs, temp)
	}

	if len(strs) > 1 {
		for _, v := range strs {
			getString5668(v, results)
		}
	} else {
		*results = append(*results, strs[0])
	}
}
