package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 141. 环形链表
 */

func Test_leetcode_141(t *testing.T) {
	node1 := &ListNode{3, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{0, nil}
	node4 := &ListNode{4, nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(hasCycle(node1))

	node := &ListNode{3, nil}
	node.Next = node
	fmt.Println(hasCycle(node))
}

func hasCycle(head *ListNode) bool {
	defer timeCost()()
	dummy := &ListNode{-1, nil}
	for head != nil {
		if head.Next != nil && (head.Next == head || head.Next.Next == dummy) {
			return true
		}
		head.Next, head = dummy, head.Next
	}
	return false
}
