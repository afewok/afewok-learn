package leetcode

import (
	"testing"
)

/**
 * 86. 分隔链表
 */

func Test_leetcode_086(t *testing.T) {
	var node *ListNode = &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	showListNode(node)
	showListNode(partition(node, 3))
}

func partition(head *ListNode, x int) *ListNode {
	defer timeCost()()
	dummy1, dummy2 := &ListNode{-1, nil}, &ListNode{-1, nil}
	p1, p2 := dummy1, dummy2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head = head.Next
	}
	p2.Next = nil
	p1.Next = dummy2.Next
	return dummy1.Next
}
