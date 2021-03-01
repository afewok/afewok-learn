package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 124. 二叉树中的最大路径和
 */

func Test_leetcode_124(t *testing.T) {
	fmt.Println(maxPathSum(&TreeNode{-5, &TreeNode{-3, &TreeNode{-1, nil, nil}, &TreeNode{-2, nil, nil}}, &TreeNode{-4, nil, nil}}))
	fmt.Println(maxPathSum(&TreeNode{0, nil, nil}))
}

func maxPathSum(root *TreeNode) int {
	defer timeCost()()
	if root != nil {
		max = root.Val
	}
	maxPathSum124(root)
	return max
}

var max int = 0

func maxPathSum124(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxPathSum124(root.Left), maxPathSum124(root.Right)
	if max < left+root.Val+right {
		max = left + root.Val + right
	}
	return max124(root.Val+max124(left, right), 0)
}

func max124(a, b int) int {
	if a > b {
		return a
	}
	return b
}
