package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 243. 最短单词距离
 */

func Test_leetcode_5634(t *testing.T) {
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
}

func shortestDistance(words []string, word1 string, word2 string) int {
	defer timeCost()()
	min, temp, w1, w2 := len(words), 0, -1, -1
	for i, word := range words {
		if word1 == word {
			w1 = i
			if w2 == -1 {
				continue
			}
			temp = w1 - w2
			if min > temp {
				min = temp
			}
		} else if word2 == word {
			w2 = i
			if w1 == -1 {
				continue
			}
			temp = w2 - w1
			if min > temp {
				min = temp
			}
		}
	}
	return min
}
