package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 107. 二叉树的层序遍历 II
 *
 */
func Test_leetcode_107(t *testing.T) {
	fmt.Println(levelOrderBottom(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
}

func levelOrderBottom(root *TreeNode) [][]int {
	defer timeCost()()
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	temp1, temp2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	temp1 = append(temp1, root)
	for len(temp1) > 0 {
		temp := make([]int, 0)
		for _, item := range temp1 {
			temp = append(temp, item.Val)
			if item.Left != nil {
				temp2 = append(temp2, item.Left)
			}
			if item.Right != nil {
				temp2 = append(temp2, item.Right)
			}
		}
		result = append([][]int{temp}, result...)
		temp1, temp2 = temp2, temp1
		temp2 = temp2[:0]
	}
	return result
}
