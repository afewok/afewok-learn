package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func showListNode(node *ListNode) {
	if node == nil {
		fmt.Println("node 为nil")
		return
	}
	for node != nil {
		fmt.Printf("%v -> ", node.Val)
		node = node.Next
	}
	fmt.Println("NULL")
}
