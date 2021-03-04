package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 剑指 Offer 26. 树的子结构
 */
func Test_leetcode_offer_026(t *testing.T) {
	fmt.Println(isSubStructure(
		&TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{5, nil, nil}},
		&TreeNode{4, &TreeNode{1, nil, nil}, nil}))
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
