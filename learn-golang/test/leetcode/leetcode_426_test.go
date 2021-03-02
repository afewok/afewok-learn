package leetcode

import (
	"testing"
)

/**
 * 426. 将二叉搜索树转化为排序的双向链表
 */

func Test_leetcode_426(t *testing.T) {
	treeNode := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}
	showTreeNode(treeNode)
	showLinkedListNode(treeToDoublyList(treeNode))
}

func treeToDoublyList(root *TreeNode) *TreeNode {
	defer timeCost()()
	head := tree426(root)
	tail := head
	for tail.Right != nil {
		tail = tail.Right
	}
	head.Left, tail.Right = tail, tail
	return head
}

func tree426(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := tree426(root.Right)
	if right != nil {
		right.Left = root
	}

	left := tree426(root.Left)
	if left != nil {
		left.Left = root
	}

	return root
}
