package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 234. 回文链表
 */

func Test_leetcode_234(t *testing.T) {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(218))
}

func isPalindrome(head *ListNode) bool {
	return isPalindrome234(head, head)
}

func isPalindrome234(head *ListNode, tail *ListNode) bool {
	if tail.Next != nil {

	}
}
