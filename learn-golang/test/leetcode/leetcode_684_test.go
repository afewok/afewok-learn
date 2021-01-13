package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 684. 冗余连接
 */

func Test_leetcode_684(t *testing.T) {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
}

// 一维数组p[i] = j表示的是第i个节点前驱节点为j, 连通树中根节点k满足p[k]==k
func findRedundantConnection(edges [][]int) []int {
	defer timeCost()()
	var res = make([]int, 2)
	var p = make([]int, len(edges)+1)
	for i := 1; i <= len(edges); i++ {
		p[i] = i
	}

	for i := 0; i < len(edges); i++ {
		a := find(edges[i][0], p)
		b := find(edges[i][1], p)
		if a == b {
			res = []int{edges[i][0], edges[i][1]}
		} else {
			// 合并
			//p[a] = b
			union(a, b, p)
		}
	}
	return res
}

// 包含了树压缩动作，将每个子集合树变为二级
func find(x int, p []int) int {
	if p[x] != x {
		p[x] = find(p[x], p)
	}

	return p[x]
}

// 对两个集合按树根节点（子集的领导节点）合并， x, y为两个集合的任意两个节点
func union(x, y int, p []int) {
	x_root, y_root := find(x, p), find(y, p)
	if x_root != y_root {
		p[x_root] = y_root
	}
}
