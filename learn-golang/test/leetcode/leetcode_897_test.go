package leetcode

import (
	"testing"
)

/**
 * 897. 递增顺序查找树
 *
 */
func Test_leetcode_897(t *testing.T) {
	showTreeNode(increasingBST(&TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}}}))
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		return &TreeNode{root.Val, nil, nil}
	}
	var result *TreeNode
	if root.Left != nil {
		result = increasingBST(root.Left)
	}
	p := result
	if result == nil {
		result = &TreeNode{root.Val, nil, nil}
		p = result
	} else {
		for p.Right != nil {
			p = p.Right
		}
		p.Right = &TreeNode{root.Val, nil, nil}
		p = p.Right
	}

	if root.Right != nil {
		temp := increasingBST(root.Right)
		if temp != nil {
			p.Right = temp
		}
	}
	return result
}
