package leetcode

import (
	"testing"
)

/**
 * 206. 反转链表
 */

func Test_leetcode_206(t *testing.T) {
	showListNode(reverseList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}))
}

func reverseList(head *ListNode) *ListNode {
	defer timeCost()()
	dummy := &ListNode{-1, nil}
	p := dummy.Next
	for head != nil {
		p, dummy.Next, head = dummy.Next, head, head.Next
		dummy.Next.Next = p
	}
	return dummy.Next
}
