package leetcode

import (
	"testing"
)

/**
 * 108. 将有序数组转换为二叉搜索树
 */

func Test_leetcode_108(t *testing.T) {
	showTreeNode(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST108(nums, 0, len(nums)-1)
}

func sortedArrayToBST108(nums []int, head, tail int) *TreeNode {
	if head > tail {
		return nil
	}
	mid := head + (tail-head)/2
	return &TreeNode{nums[mid], sortedArrayToBST108(nums, head, mid-1), sortedArrayToBST108(nums, mid+1, tail)}
}
