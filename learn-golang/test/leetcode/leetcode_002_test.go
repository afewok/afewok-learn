package leetcode

import "testing"

/**
 * 2. 两数相加
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例： 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4) 输出：7 -> 0 -> 8 原因：342 + 465 = 807
 *
 * 思路：虚拟头结点遍历、递归维护一个节点
 */

func Test_leetcode_002(t *testing.T) {
	arrayPoint1 := preData()
	showListNode(arrayPoint1[0])
	showListNode(arrayPoint1[1])
	result1 := addTwoNumbers1(arrayPoint1[0], arrayPoint1[1])
	showListNode(result1)

	arrayPoint2 := preData()
	showListNode(arrayPoint2[0])
	showListNode(arrayPoint2[1])
	result2 := addTwoNumbers2(arrayPoint2[0], arrayPoint2[1])
	showListNode(result2)
}

func preData() []*ListNode {
	var l1 *ListNode = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	var l2 *ListNode = &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	var result []*ListNode = make([]*ListNode, 2)
	result[0] = l1
	result[1] = l2
	return result
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var addOne bool = false
	var dummy *ListNode = &ListNode{0, nil}
	temp := dummy

	for addOne || l1 != nil || l2 != nil {
		switch {
		case addOne:
			temp.Next = &ListNode{1, nil}
		default:
			temp.Next = &ListNode{0, nil}
		}

		temp = temp.Next
		if l1 != nil {
			temp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp.Val += l2.Val
			l2 = l2.Next
		}
		switch {
		case temp.Val > 9:
			addOne = true
		default:
			addOne = false
		}

		temp.Val = temp.Val % 10
	}
	return dummy.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	val := 0
	var addOne bool = false
	var l2Next *ListNode = nil
	if l2 != nil && l2.Next != nil {
		l2Next = l2.Next
	}
	if l2 != nil {
		val += l2.Val
	}

	l1.Val += val
	if l1.Val > 9 {
		addOne = true
	}
	l1.Val %= 10

	if l1.Next == nil && (addOne || l2Next != nil) {
		l1.Next = new(ListNode)
	}

	l1Next := l1.Next
	if addOne {
		l1Next.Val++
	}

	if l1Next != nil || l2Next != nil {
		addTwoNumbers2(l1Next, l2Next)
	}
	return l1
}
