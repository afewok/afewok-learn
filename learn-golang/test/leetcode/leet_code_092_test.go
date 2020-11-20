package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 92. 反转链表 II
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。 说明: 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例: 输入: 1->2->3->4->5->NULL, m = 2, n = 4 输出: 1->4->3->2->5->NULL
 *
 * 思路：递归、虚拟头结点(迭代链接反转)
 */
func Test_leet_code_092(t *testing.T) {
	m, n := 2, 4
	var head1 *ListNode = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	show(head1)
	result1 := reverseBetween1(head1, m, n)
	show(result1)
}

func show(head *ListNode) {
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Println("NULL")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	virtual := &ListNode{0, head}
	pre := virtual
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	head = pre.Next
	for i := m; i < n; i++ {
		//P ->      H -> T -> S
		//1 -> 3 -> 2 -> 4 -> 5
		temp := head.Next
		head.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}

	return virtual.Next
}
