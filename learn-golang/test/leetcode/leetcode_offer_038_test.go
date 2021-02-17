package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 剑指 Offer 38. 字符串的排列
 * 思路：深度优先遍历（递归+回溯+剪枝）、广度优先遍历（动态规划（杨辉三角第n行））
 */
func Test_leetcode_offer_038(t *testing.T) {
	fmt.Println(permutation("abc"))
}

func permutation1(s string) []string {
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

func permutation(s string) []string {
	defer timeCost()()
	result := make([]string, 0)
	getPermutation(s, string(s[0]), 1, len(s), make(map[string]int), &result)
	return result

}

func getPermutation(s, temp string, sub, length int, mp map[string]int, result *[]string) {
	if sub == length {
		if _, OK := mp[temp]; !OK {
			mp[temp]++
			*result = append(*result, temp)
		}
		return
	}

	for j := 0; j <= len(temp); j++ {
		getPermutation(s, temp[:j]+string(s[sub])+temp[j:], sub+1, length, mp, result)
	}
}
