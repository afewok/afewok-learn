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
	defer timeCost()()
	return true
}
