package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 404. 左叶子之和
 */

func Test_leetcode_404(t *testing.T) {
	fmt.Println(sumOfLeftLeaves(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	return getSumOfLeftLeaves(root)
}
func getSumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return root.Val
	}
	result := getSumOfLeftLeaves(root.Left)
	if root.Right != nil && (root.Right.Left != nil || root.Right.Right != nil) {
		result += getSumOfLeftLeaves(root.Right)
	}
	return result
}
