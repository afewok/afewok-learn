package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1002. 查找常用字符
 */

func Test_leetcode_1002(t *testing.T) {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}

func commonChars(A []string) []string {
	defer timeCost()()
	mp, tempMp, length, result := make(map[rune]int), make(map[rune]int), len(A), make([]string, 0)
	for _, c := range A[0] {
		mp[c]++
	}
	for i := 1; i < length; i++ {
		for _, c := range A[i] {
			tempMp[c]++
		}
		for k, count1 := range mp {
			count2, OK := tempMp[k]
			if !OK {
				delete(mp, k)
			} else if count1 < count2 {
				mp[k] = count1
			} else if count1 > count2 {
				mp[k] = count2
			}
		}
		for k := range tempMp {
			delete(tempMp, k)
		}
	}

	for key, count := range mp {
		for count > 0 {
			count--
			result = append(result, string(key))
		}
	}
	return result
}
