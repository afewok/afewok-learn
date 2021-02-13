package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 94. 二叉树的中序遍历
 */

func Test_leetcode_094(t *testing.T) {
	fmt.Println(inorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}))
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}
