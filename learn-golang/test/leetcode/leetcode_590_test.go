package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 590. N 叉树的后序遍历
 */

func Test_leetcode_590(t *testing.T) {
	fmt.Println(postorder(&Node{1, []*Node{&Node{3, []*Node{&Node{5, nil}, &Node{6, nil}}}, &Node{2, nil}, &Node{4, nil}}}))
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	for _, children := range root.Children {
		temp := postorder(children)
		if temp != nil {
			result = append(result, temp...)
		}
	}
	result = append(result, root.Val)
	return result
}
