package leetcode

import (
	"testing"
)

/**
 * 23. 合并K个升序链表
 */

func Test_leetcode_023(t *testing.T) {
	showListNode(mergeKLists([]*ListNode{}))
	showListNode(mergeKLists([]*ListNode{{}}))
	showListNode(mergeKLists([]*ListNode{&ListNode{1, &ListNode{4, &ListNode{5, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}, &ListNode{2, &ListNode{6, nil}}}))
}

func mergeKLists(lists []*ListNode) *ListNode {
	defer timeCost()()
	length := len(lists)
	if length == 0 {
		return nil
	}

	dummy := &ListNode{-1, lists[0]}
	for i := 1; i < length; i++ {
		p := dummy
		head := lists[i]
		for head != nil {
			if p.Next == nil {
				p.Next = head
				break
			} else if p.Next.Val > head.Val {
				temp := head.Next
				p.Next, head.Next, head = head, p.Next, temp
			}
			p = p.Next
		}
	}
	return dummy.Next
}
