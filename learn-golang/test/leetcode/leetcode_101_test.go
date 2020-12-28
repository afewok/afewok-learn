package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 101. 对称二叉树
 */

func Test_leetcode_101(t *testing.T) {
	fmt.Println(isSymmetric(&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}))
}

func isSymmetric(root *TreeNode) bool {
	defer timeCost()()
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	} else if !check(left.Left, right.Right) {
		return false
	}
	return check(left.Right, right.Left)
}
