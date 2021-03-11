package leetcode

import (
	"testing"
)

/**
 * 147. 对链表进行插入排序
 */

func Test_leetcode_147(t *testing.T) {
	showListNode(insertionSortList(&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}))
}

func insertionSortList(head *ListNode) *ListNode {
	defer timeCost()()
	dummy := &ListNode{-1, head}
	for head != nil && head.Next != nil {
		if head.Val <= head.Next.Val {
			head = head.Next
			continue
		}
		temp, prev := head.Next, dummy
		head.Next = head.Next.Next
		for prev.Next.Val < temp.Val {
			prev = prev.Next
		}
		temp.Next, prev.Next = prev.Next, temp
	}
	return dummy.Next
}
