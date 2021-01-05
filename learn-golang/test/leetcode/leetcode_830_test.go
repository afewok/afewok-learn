package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 830. 较大分组的位置
 */

func Test_leetcode_830(t *testing.T) {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("abc"))
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
	fmt.Println(largeGroupPositions("aba"))
	fmt.Println(largeGroupPositions("aaa"))
}

func largeGroupPositions(s string) [][]int {
	defer timeCost()()
	arr, sub, length := make([][]int, 0), 0, len(s)
	for i := 1; i < length; i++ {
		if s[sub] != s[i] {
			if i-sub >= 3 {
				arr = append(arr, []int{sub, i - 1})
			}
			sub = i
		}
	}
	if length-sub >= 3 {
		arr = append(arr, []int{sub, length - 1})
	}
	return arr
}
