package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 606. 根据二叉树创建字符串
 */

func Test_leetcode_606(t *testing.T) {
	fmt.Println(tree2str(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}}))
	fmt.Println(tree2str(&TreeNode{1, &TreeNode{2, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}))
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	left := tree2str(t.Left)
	right := tree2str(t.Right)
	if left == "" && right == "" {
		return strconv.Itoa(t.Val)
	}else if right == ""{
		return strconv.Itoa(t.Val) + "(" + left + ")"
	}
	return strconv.Itoa(t.Val) + "(" + left + ")" + "(" + right + ")"
}
