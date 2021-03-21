package leetcode

import "testing"

/**
 * 82. 删除排序链表中的重复元素 II
 */

func Test_leetcode_082(t *testing.T) {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	showListNode(node)
	showListNode(deleteDuplicates082(node))
}

func deleteDuplicates082(head *ListNode) *ListNode {
	defer timeCost()()
	if head == nil || head.Next == nil {
		return head
	}
	dummy, mp := &ListNode{-1, head}, make(map[int]int)
	for head != nil {
		mp[head.Val]++
		head = head.Next
	}
	p, q := dummy.Next, dummy
	for p != nil {
		if v, _ := mp[p.Val]; v > 1 {
			q.Next = p.Next
		} else {
			q = p
		}
		p = p.Next
	}
	return dummy.Next
}
