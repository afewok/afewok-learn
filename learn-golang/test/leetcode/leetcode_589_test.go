package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 589. N 叉树的前序遍历
 */

func Test_leetcode_589(t *testing.T) {
	fmt.Println(preorder(&Node{1, []*Node{&Node{3, []*Node{&Node{5, nil}, &Node{6, nil}}}, &Node{2, nil}, &Node{4, nil}}}))
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	for _, children := range root.Children {
		temp := preorder(children)
		if temp != nil {
			result = append(result, temp...)
		}
	}
	return result
}
