package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 剑指 Offer 38. 字符串的排列
 */
func Test_leetcode_offer_038(t *testing.T) {
	fmt.Println(permutation("abc"))
}

func permutation(s string) []string {
	defer timeCost()()
	length := len(s)
	mp, pre, result := make(map[string]int), make([]string, 0), []string{string(s[0])}
	for i := 1; i < length; i++ {
		for _, str := range result {
			ll := len(str)
			for j := 0; j <= ll; j++ {
				temp := str[:j] + string(s[i]) + str[j:]
				if _, OK := mp[temp]; !OK {
					mp[temp]++
					pre = append(pre, temp)
				}
			}
		}
		result, pre = pre, make([]string, 0)
	}
	return result
}
