package leetcode

import (
	"testing"
)

/**
 * 25. K 个一组翻转链表
 *
 */
func Test_leetcode_025(t *testing.T) {
	showListNode(reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 1))
	showListNode(reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
	showListNode(reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 3))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	defer timeCost()()
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{-1, nil}
	d, p := dummy, head
	for head != nil {
		for i := 0; i < k; i++ {
			if p == nil {
				d.Next = head
				return dummy.Next
			}
			p = p.Next
		}
		q := head
		for head != p {
			temp := head.Next
			head.Next, d.Next, head = d.Next, head, temp
		}
		d = q
	}
	return dummy.Next
}
