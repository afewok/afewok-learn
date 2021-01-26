package leetcode

import (
	"testing"
)

/**
 * 876. 链表的中间结点
 *
 * 思路：
 */
func Test_leetcode_876(t *testing.T) {
	showListNode(middleNode(nil))
	showListNode(middleNode(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}))
	showListNode(middleNode(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}))
}

func middleNode(head *ListNode) *ListNode {
	defer timeCost()()
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	return list[len(list)/2]
}
