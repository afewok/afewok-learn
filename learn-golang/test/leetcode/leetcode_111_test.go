package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 111. 二叉树的最小深度
 */

func Test_leetcode_111(t *testing.T) {
	node := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	showTreeNode(node)
	fmt.Println(minDepth(node))
	node = &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}}}
	showTreeNode(node)
	fmt.Println(minDepth(node))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	temp1, temp2 := minDepth(root.Left), minDepth(root.Right)
	if temp1 == 0 {
		return temp2 + 1
	} else if temp2 == 0 {
		return temp1 + 1
	} else if temp1 < temp2 {
		return temp1 + 1
	}
	return temp2 + 1
}
