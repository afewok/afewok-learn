package leetcode

import (
	"fmt"
	"time"
)

/**
 * dp、dfs、bfs、剪枝、贪心、回溯
 * DFS 深度优先搜索、.BFS 广度优先搜索、DP算法（Dynamic Programming,俗称动态规划）
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}
