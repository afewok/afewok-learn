package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 226. 翻转二叉树
 */
func Test_leetcode_226(t *testing.T) {
	node := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}
	fmt.Println(invertTree(node))
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
