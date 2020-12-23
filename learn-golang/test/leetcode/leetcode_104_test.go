package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 104. 二叉树的最大深度
 */

func Test_leetcode_104(t *testing.T) {
	node := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(maxDepth(node))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left), maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
