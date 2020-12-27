package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5637. 判断字符串的两半是否相似
 */

func Test_leetcode_5637(t *testing.T) {
	fmt.Println(halvesAreAlike("book"))
	fmt.Println(halvesAreAlike("textbook"))
	fmt.Println(halvesAreAlike("MerryChristmas"))
	fmt.Println(halvesAreAlike("AbCdEfGh"))
}

func halvesAreAlike(s string) bool {
	defer timeCost()()
	mp, length, a, b := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}, len(s), 0, 0
	for i := 0; i < length/2; i++ {
		if _, OK := mp[s[i]]; OK {
			a++
		}
	}
	for i := length / 2; i < length; i++ {
		if _, OK := mp[s[i]]; OK {
			b++
		}
	}
	return a == b
}
