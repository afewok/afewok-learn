package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 144. 二叉树的前序遍历
 */

func Test_leetcode_144(t *testing.T) {
	fmt.Println(preorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}))
	fmt.Println(preorderTraversal(nil))
	fmt.Println(preorderTraversal(&TreeNode{1, nil, nil}))
	fmt.Println(preorderTraversal(&TreeNode{1, &TreeNode{2, nil, nil}, nil}))
	fmt.Println(preorderTraversal(&TreeNode{1, nil, &TreeNode{2, nil, nil}}))
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	list := []int{root.Val}
	list = append(list, preorderTraversal(root.Left)...)
	list = append(list, preorderTraversal(root.Right)...)
	return list
}
