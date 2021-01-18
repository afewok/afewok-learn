package leetcode

import (
	"testing"
)

/**
 * 617. 合并二叉树
 */

func Test_leetcode_617(t *testing.T) {
	showTreeNode(mergeTrees(&TreeNode{1, &TreeNode{3, &TreeNode{5, nil, nil}, nil}, &TreeNode{2, nil, nil}},
		&TreeNode{2, &TreeNode{1, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, &TreeNode{7, nil, nil}}}))
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	defer timeCost()()
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	dummy := &TreeNode{-1, nil, nil}
	mergeTree(dummy, t1, t2)
	return dummy
}

func mergeTree(t *TreeNode, t1 *TreeNode, t2 *TreeNode) {
	t.Val = t1.Val + t2.Val

	if t1.Left != nil && t2.Left == nil {
		t.Left = t1.Left
	} else if t1.Left == nil && t2.Left != nil {
		t.Left = t2.Left
	} else if t1.Left != nil && t2.Left != nil {
		t.Left = &TreeNode{-1, nil, nil}
		mergeTree(t.Left, t1.Left, t2.Left)
	}

	if t1.Right != nil && t2.Right == nil {
		t.Right = t1.Right
	} else if t1.Right == nil && t2.Right != nil {
		t.Right = t2.Right
	} else if t1.Right != nil && t2.Right != nil {
		t.Right = &TreeNode{-1, nil, nil}
		mergeTree(t.Right, t1.Right, t2.Right)
	}
}
