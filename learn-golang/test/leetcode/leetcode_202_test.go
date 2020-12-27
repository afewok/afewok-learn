package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 202. 快乐数
 */

func Test_leetcode_202(t *testing.T) {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	defer timeCost()()
	mp := map[int]bool{}
	_, OK := mp[n]
	for !OK {
		mp[n] = true
		str := strconv.Itoa(n)
		temp := 0
		for _, c := range str {
			i := int(c - 48)
			temp = temp + i*i
		}
		if temp == 1 {
			return true
		}
		n = temp
		_, OK = mp[n]
	}
	return false
}
