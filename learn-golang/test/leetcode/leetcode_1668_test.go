package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1668. 最大重复子字符串
 */

func Test_leetcode_1668(t *testing.T) {
	fmt.Println(maxRepeating("ababc", "ab"))
	fmt.Println(maxRepeating("ababc", "ba"))
	fmt.Println(maxRepeating("ababc", "ac"))
	fmt.Println(maxRepeating("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"))
}

func maxRepeating(sequence string, word string) int {
	defer timeCost()()
	len1, len2, temp, max := len(sequence), len(word), 0, 0
	for i := 0; i < len1; i++ {
		if sequence[i] != word[0] {
			continue
		}
		temp = 0
		for m, sub := i, 0; m < len1; m++ {
			if sequence[m] != word[sub] {
				break
			}
			sub++
			if sub == len2 {
				temp++
				sub = 0
			}
		}
		if max < temp {
			max = temp
		}
	}
	return max
}
