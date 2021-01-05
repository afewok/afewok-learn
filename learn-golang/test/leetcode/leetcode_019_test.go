package leetcode

import (
	"testing"
)

/**
 * 19. 删除链表的倒数第N个节点
 */

func Test_leetcode_019(t *testing.T) {
	showListNode(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	defer timeCost()()
	dummy, length := &ListNode{-1, head}, 0
	for head != nil {
		length++
		head = head.Next
	}
	length = length - n + 1
	head = dummy
	for head != nil {
		length--
		if length == 0 {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
	return dummy.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	defer timeCost()()
	if n == 0 {
		return head
	}
	dummy := &ListNode{-1, head}
	list := make([]*ListNode, 0)
	list = append(list, dummy)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	sub := len(list) - n
	if n == 1 {
		list[sub-1].Next = nil
	} else {
		list[sub-1].Next = list[sub+1]
	}
	return dummy.Next
}
