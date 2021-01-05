package leetcode

import (
	"testing"
)

/**
 * 24. 两两交换链表中的节点
 */

func Test_leetcode_024(t *testing.T) {
	showListNode(reversal(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
	showListNode(swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
	showListNode(swapPairs(nil))
}

func swapPairs(head *ListNode) *ListNode {
	defer timeCost()()
	dummy := &ListNode{-1, head}
	p := dummy
	for head != nil && head.Next != nil {
		p.Next = head.Next
		head.Next = head.Next.Next
		p.Next.Next = head

		p = p.Next.Next
		head = p.Next
	}
	return dummy.Next
}

func reversal(head *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p := dummy
	for head != nil {
		p, dummy.Next = dummy.Next, head
		head, dummy.Next.Next = dummy.Next.Next, p
	}
	return dummy.Next
}
