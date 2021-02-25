package leetcode

import (
	"testing"
)

/**
 * 160. 相交链表
 */

func Test_leetcode_160(t *testing.T) {
	node := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	showListNode(getIntersectionNode(&ListNode{4, &ListNode{1, node}}, &ListNode{5, &ListNode{0, &ListNode{1, node}}}))
	node = &ListNode{2, &ListNode{4, nil}}
	showListNode(getIntersectionNode(&ListNode{0, &ListNode{9, &ListNode{1, node}}}, &ListNode{3, node}))
	showListNode(getIntersectionNode(&ListNode{2, &ListNode{6, &ListNode{4, nil}}}, &ListNode{1, &ListNode{5, nil}}))
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	defer timeCost()()
	mp := make(map[*ListNode]bool)
	for headA != nil {
		mp[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if _, OK := mp[headB]; OK {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := 0
	tmpA := headA
	for tmpA != nil {
		lengthA++
		tmpA = tmpA.Next
	}

	lengthB := 0
	tmpB := headB
	for tmpB != nil {
		lengthB++
		tmpB = tmpB.Next
	}

	if lengthA > lengthB {
		for lengthA != lengthB {
			headA = headA.Next
			lengthA--
		}
	} else {
		for lengthB != lengthA {
			headB = headB.Next
			lengthB--
		}
	}

	// 用headA或者headB做判断
	for headA != nil {
		if headA == headB {
			return headA
		} else {
			headA = headA.Next
			headB = headB.Next
		}
	}

	return nil
}
