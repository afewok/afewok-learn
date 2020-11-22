package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 3. 无重复字符的最长子串
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:输入: "abcabcbb" 输出: 3 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 * 示例 2:输入: "bbbbb" 输出: 1 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 * 示例 3:输入: "pwwkew" 输出: 3 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。 请注意，你的答案必须是子串
 * 的长度，"pwke" 是一个子序列，不是子串。
 *
 * 思路：双指针+滑动窗口、动态规划+哈希表、利用数组（桶）代替hashmap、暴力双循环
 */
func Test_leetcode_003(t *testing.T) {
	s1 := "abcabcbb"
	fmt.Println(s1, lengthOfLongestSubstring(s1))
	s2 := "bbbbb"
	fmt.Println(s2, lengthOfLongestSubstring(s2))
}

func lengthOfLongestSubstring(s string) int {
	var max int = 0
	return max
}
