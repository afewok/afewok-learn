package leetcode

import (
	"testing"
)

/**
 * 61. 旋转链表
 */

func Test_leetcode_061(t *testing.T) {
	// showListNode(rotateRight(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
	// showListNode(rotateRight(&ListNode{0, &ListNode{1, &ListNode{2, nil}}}, 4))
	showListNode(rotateRight(&ListNode{1, &ListNode{2, nil}}, 2))
	// showListNode(rotateRight(&ListNode{1, nil}, 1))
}

func rotateRight(head *ListNode, k int) *ListNode {
	defer timeCost()()
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	count, dummy, tail, q := 0, &ListNode{-1, head}, head, head
	for head != nil {
		count++
		tail = head
		head = head.Next
	}
	head = dummy.Next
	k = k % count
	for i := 0; i < count-k; i++ {
		q = head
		head = head.Next
	}
	if head == nil {
		return dummy.Next
	}
	tail.Next, q.Next = dummy.Next, nil
	return head
}
