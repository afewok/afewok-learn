package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 234. 回文链表
 */

func Test_leetcode_234(t *testing.T) {
	fmt.Println(isPalindrome(&ListNode{1, &ListNode{2, nil}}))
	fmt.Println(isPalindrome(&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}))
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	temp = head
	return isPalindrome234(head)
}

var temp *ListNode

func isPalindrome234(tail *ListNode) bool {
	if tail.Next != nil {
		if !isPalindrome234(tail.Next) {
			return false
		}
	}
	if temp.Val != tail.Val {
		return false
	}
	temp = temp.Next
	return true
}
