package leetcode

import (
	"testing"
)

/**
 * 1008. 前序遍历构造二叉搜索树
 */

func Test_leetcode_1008(t *testing.T) {
	showTreeNode(bstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
}

func bstFromPreorder(preorder []int) *TreeNode {
	defer timeCost()()
	length := len(preorder)
	node := TreeNode{preorder[0], nil, nil}
	for i := 1; i < length; i++ {
		getBstFromPreorder(&node, preorder[i])
	}
	return &node
}

func getBstFromPreorder(node *TreeNode, num int) {
	if num < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{num, nil, nil}
		} else {
			getBstFromPreorder(node.Left, num)
		}
	}
	if num > node.Val {
		if node.Right == nil {
			node.Right = &TreeNode{num, nil, nil}
		} else {
			getBstFromPreorder(node.Right, num)
		}
	}
}
