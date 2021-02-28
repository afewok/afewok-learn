package leetcode

import (
	"fmt"
	"time"
)

/**
 * dp、dfs、bfs、剪枝、贪心、回溯
 * DFS 深度优先搜索、BFS 广度优先搜索、DP算法（Dynamic Programming,俗称动态规划）
 */
func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func showTreeNode(node *TreeNode) {
	if node == nil {
		fmt.Println("node 为nil")
		return
	}
	getShowTreeNode(node, "", true)
}

func getShowTreeNode(node *TreeNode, prefix string, isLeft bool) {
	if isLeft {
		if node.Right != nil {
			getShowTreeNode(node.Right, prefix+"│   ", false)
		}
		fmt.Printf("%v%v%d\n", prefix, "└── ", node.Val)
		if node.Left != nil {
			getShowTreeNode(node.Left, prefix+"    ", true)
		}
	} else {
		if node.Right != nil {
			getShowTreeNode(node.Right, prefix+"    ", false)
		}
		fmt.Printf("%v%v%d\n", prefix, "┌── ", node.Val)
		if node.Left != nil {
			getShowTreeNode(node.Left, prefix+"│   ", true)
		}
	}
}

func showLinkedListNode(root *TreeNode) {
	if root == nil {
		fmt.Println("node 为nil")
		return
	}
	p := root
	for {
		fmt.Printf("%v <=> ", p.Val)
		p = p.Right
		if p == nil || p.Left == nil || p.Left.Right == nil {
			fmt.Println("NULL")
			break
		} else if p.Left.Right != p {
			fmt.Println("No Double Linked List")
			break
		} else if p == root {
			break
		}
	}
	fmt.Println("")
}

func maxTwo(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxThree(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}
