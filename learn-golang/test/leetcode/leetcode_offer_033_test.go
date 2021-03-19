package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 剑指 Offer 33. 二叉搜索树的后序遍历序列
 */
func Test_leetcode_offer_033(t *testing.T) {
	// fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
	// fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
	fmt.Println(verifyPostorder([]int{5, 2, -17, -11, 25, 76, 62, 98, 92, 61}))
}

func verifyPostorder(postorder []int) bool {
	return verifyPostorder33(postorder, 0, len(postorder)-1)
}

func verifyPostorder33(postorder []int, start, tail int) bool {
	temp := tail - start
	if temp >= 2 {
		for i := start; i < tail; i++ {
			if postorder[i] < postorder[tail] {
				continue
			}
			for j := i; j < tail; j++ {
				if postorder[j] <= postorder[tail] {
					return false
				}
			}
			return verifyPostorder33(postorder, start, i-1) && verifyPostorder33(postorder, i, tail-1)
		}
		return verifyPostorder33(postorder, start, tail-1)
	}

	return true
}
