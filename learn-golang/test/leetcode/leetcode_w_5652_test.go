package leetcode

import (
	"testing"
)

/**
 * 5652. 交换链表中的节点
 */

func Test_leetcode_5652(t *testing.T) {
	showListNode(swapNodes(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
}

func swapNodes(head *ListNode, k int) *ListNode {
	defer timeCost()()
	if head == nil {
		return nil
	}
	dummy, arr := &ListNode{-1, head}, make([]*ListNode, 0)
	arr = append(arr, dummy)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	length := len(arr)
	arr[k-1].Next, arr[length-k-1].Next = arr[length-k], arr[k]
	arr[length-k].Next, arr[k].Next = arr[k].Next, arr[length-k].Next
	arr[k], arr[length-k] = arr[length-k], arr[k]
	return dummy.Next
}
