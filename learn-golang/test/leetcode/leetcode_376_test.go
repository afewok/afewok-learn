package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 376. 摆动序列
 *
 * 思路：
 */
func Test_leetcode_376(t *testing.T) {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

type Node struct {
}

func wiggleMaxLength(nums []int) int {
	defer timeCost()()
	if nums == nil {
		return 0
	}
	var length int = len(nums)
	if length == 0 {
		return 0
	}
	i, cSub, clen, pSub, plen, mSub, mlen := 0, 0, 1, 0, 0, 0, 0
	flag := true
	for i = 1; i < length; i++ {
		if nums[i-1] > nums[i] {

		} else if nums[i-1] < nums[i] {

		}
	}
	//最长结尾需要负
	//最长结尾需要正
	//当前长度
	for i = 1; i < length; i++ {

		//结尾，下一步，长度
	}

	return clen
}
