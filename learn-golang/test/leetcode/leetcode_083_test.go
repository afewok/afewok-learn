package leetcode

import "testing"

/**
 * 83. 删除排序链表中的重复元素
 */

func Test_leetcode_083(t *testing.T) {
	node := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	showListNode(node)
	showListNode(deleteDuplicates(node))
}

func deleteDuplicates(head *ListNode) *ListNode {
	defer timeCost()()
	if head == nil {
		return head
	}
	dummy, p := &ListNode{0, head}, head.Next
	for p != nil {
		if head.Val == p.Val {
			head.Next = p.Next
		} else {
			head = head.Next
		}
		p = head.Next
	}
	return dummy.Next
}
