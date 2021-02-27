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

	return root
}

// func tree426(root *TreeNode) (*TreeNode,*TreeNode){

// }
