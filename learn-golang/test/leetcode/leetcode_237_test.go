package leetcode

import (
	"testing"
)

/**
 * 237. 删除链表中的节点
 */

func Test_leetcode_237(t *testing.T) {
	list := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}
	deleteNode(list.Next)
	showListNode(list)
}

func deleteNode(node *ListNode) {
	defer timeCost()()
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil
}

func deleteNode1(node *ListNode) {
	defer timeCost()()
	for node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			break
		}
		node = node.Next
	}
}
