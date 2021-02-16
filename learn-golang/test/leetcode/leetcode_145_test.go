package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 145. 二叉树的后序遍历
 */

func Test_leetcode_145(t *testing.T) {
	fmt.Println(postorderTraversal(&TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}))
	fmt.Println(postorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}))
	fmt.Println(postorderTraversal(nil))
	fmt.Println(postorderTraversal(&TreeNode{1, nil, nil}))
	fmt.Println(postorderTraversal(&TreeNode{1, &TreeNode{2, nil, nil}, nil}))
	fmt.Println(postorderTraversal(&TreeNode{1, nil, &TreeNode{2, nil, nil}}))
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	list := postorderTraversal(root.Left)
	list = append(list, postorderTraversal(root.Right)...)
	list = append(list, root.Val)
	return list
}
