package leetcode

import (
	"testing"
)

/**
 * 889. 根据前序（根左右）和后序（左右根）遍历构造二叉树
 *
 */
func Test_leetcode_889(t *testing.T) {
	showTreeNode(constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}))
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	length := len(pre)
	if length == 0 {
		return nil
	}
	node := TreeNode{pre[0], nil, nil}
	if length == 1 {
		return &node
	}
	for i := range post {
		if post[i] == pre[1] {
			node.Left = constructFromPrePost(pre[1:i+2], post[0:i+1])
			node.Right = constructFromPrePost(pre[i+2:length], post[i+1:length-1])
			break
		}
	}
	return &node
}

func constructFromPrePost1(pre []int, post []int) *TreeNode {
	length := len(pre)
	if length == 0 {
		return nil
	}
	node := &TreeNode{pre[0], nil, nil}
	if length == 1 {
		return node
	}
	for i := length - 1; i > 0; i-- {
		if post[i-1] == pre[1] {
			node.Left = constructFromPrePost(pre[1:i+1], post[0:i])
			node.Right = constructFromPrePost(pre[i+1:], post[i:length-1])
			break
		}
	}
	return node
}
