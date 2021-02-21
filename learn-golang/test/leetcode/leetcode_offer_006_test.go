package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 剑指 Offer 06. 从尾到头打印链表
 */
func Test_leetcode_offer_006(t *testing.T) {
	fmt.Println(reversePrint(&ListNode{1, &ListNode{3, &ListNode{2, nil}}}))
}

func reversePrint(head *ListNode) []int {
	defer timeCost()()
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	length := len(result)
	count := (length + 1) / 2
	for i := 0; i < count; i++ {
		result[i], result[length-i-1] = result[length-i-1], result[i]
	}
	return result
}
