package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 500. 键盘行
 */

func Test_leetcode_500(t *testing.T) {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

func findWords(words []string) []string {
	defer timeCost()()
	result := make([]string, 0)
	mp1 := map[byte]bool{'Q': true, 'W': true, 'E': true, 'R': true, 'T': true, 'Y': true, 'U': true, 'I': true, 'O': true, 'P': true, 'q': true, 'w': true, 'e': true, 'r': true, 't': true, 'y': true, 'u': true, 'i': true, 'o': true, 'p': true}
	mp2 := map[byte]bool{'A': true, 'S': true, 'D': true, 'F': true, 'G': true, 'H': true, 'J': true, 'K': true, 'L': true, 'a': true, 's': true, 'd': true, 'f': true, 'g': true, 'h': true, 'j': true, 'k': true, 'l': true}
	mp3 := map[byte]bool{'Z': true, 'X': true, 'C': true, 'V': true, 'B': true, 'N': true, 'M': true, 'z': true, 'x': true, 'c': true, 'v': true, 'b': true, 'n': true, 'm': true}
	for _, word := range words {
		var mp map[byte]bool
		exist, length := true, len(word)
		if length < 1 {
			continue
		}
		if _, OK := mp1[word[0]]; OK {
			mp = mp1
		} else if _, OK := mp2[word[0]]; OK {
			mp = mp2
		} else if _, OK := mp3[word[0]]; OK {
			mp = mp3
		}
		for i := 1; i < length; i++ {
			if _, OK := mp[word[i]]; !OK {
				exist = false
				break
			}
		}
		if exist {
			result = append(result, word)
		}
	}
	return result
}
