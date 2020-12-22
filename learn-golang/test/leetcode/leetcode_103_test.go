package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 103. 二叉树的锯齿形层序遍历
 */

func Test_leetcode_103(t *testing.T) {
	node := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(zigzagLevelOrder(node))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	dfs(root, 0, &ans)
	return ans
}

func dfs(root *TreeNode, depth int, ans *[][]int) {
	if root == nil {
		return
	}
	if depth >= len(*ans) {
		*ans = append(*ans, []int{})
	}
	if depth&1 == 0 {
		(*ans)[depth] = append((*ans)[depth], root.Val)
	} else {
		// expensive operation
		(*ans)[depth] = append([]int{root.Val}, (*ans)[depth]...)
	}
	dfs(root.Left, depth+1, ans)
	dfs(root.Right, depth+1, ans)
}
