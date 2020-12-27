package leetcode

import (
	"testing"
)

/**
 * 203. 移除链表元素
 */

func Test_leetcode_203(t *testing.T) {
	node := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	showListNode(removeElements(node, 6))
}

func removeElements(head *ListNode, val int) *ListNode {
	defer timeCost()()
	dummy := &ListNode{-1, head}
	head = dummy
	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
