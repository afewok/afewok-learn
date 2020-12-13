package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 49. 字母异位词分组
 *
 * 思路：
 */
func Test_leetcode_649(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	defer timeCost()()
	mp := make(map[int]string)
	for _, str := range strs {
		count := 0
		for _, ch := range str {
			count = count + int(ch)
		}
		if v, OK := mp[count]; OK {

		}
	}
}
