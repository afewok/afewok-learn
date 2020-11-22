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
 * 思路：双指针+滑动窗口、动态规划+哈希表（虚拟-1下标）、利用数组（桶）代替hashmap、暴力双循环.....其核心是记左下标，还是记上一次长度
 */
func Test_leetcode_003(t *testing.T) {
	s1 := "abcabcbb"
	fmt.Println(s1, lengthOfLongestSubstring1(s1))
	fmt.Println(s1, lengthOfLongestSubstring2(s1))
	fmt.Println(s1, lengthOfLongestSubstring3(s1))
	s2 := "pwwkew"
	fmt.Println(s2, lengthOfLongestSubstring1(s2))
	fmt.Println(s2, lengthOfLongestSubstring2(s2))
	fmt.Println(s2, lengthOfLongestSubstring3(s2))
	s3 := "bbbbb"
	fmt.Println(s3, lengthOfLongestSubstring1(s3))
	fmt.Println(s3, lengthOfLongestSubstring2(s3))
	fmt.Println(s3, lengthOfLongestSubstring3(s3))
}

func lengthOfLongestSubstring1(s string) int {
	var sub, max, length int = -1, 0, len(s)
	var mapData map[byte]int = make(map[byte]int, length)
	for i := 0; i < length; i++ {
		c := s[i]
		if v, ok := mapData[c]; ok && sub < v {
			sub = v
		}
		if v := i - sub; max < v {
			max = v
		}
		mapData[c] = i
	}
	return max
}

func lengthOfLongestSubstring2(s string) int {
	var count, max, length int = 0, 0, len(s)
	mapData := map[byte]int{}
	for i := 0; i < length; i++ {
		c := s[i]
		if v, ok := mapData[c]; ok && count+1 > i-v {
			count = i - v
		} else {
			count++
		}
		if max < count {
			max = count
		}
		mapData[c] = i
	}
	return max
}

func lengthOfLongestSubstring3(s string) int {
	var rSub, max, length int = 0, 0, len(s)
	mapData := map[byte]int{}
	for i := 0; i < length; i++ {
		if i > 0 {
			delete(mapData, s[i-1])
		}

		for rSub < length && mapData[s[rSub]] == 0 { //根据map的这个value的类型，无key时返回缺省值
			mapData[s[rSub]]++
			rSub++
			if v := rSub - i; max < v {
				max = v
			}
		}

		if max > length-i {
			break
		}
	}
	return max
}
