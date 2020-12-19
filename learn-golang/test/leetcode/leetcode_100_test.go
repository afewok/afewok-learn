package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 100. 相同的树
 */
func Test_leetcode_100(t *testing.T) {
	node1, node2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fc := timeCost()
	fmt.Println(isSameTree(node1, node2))
	fc()
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil {
		return false
	} else if q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
