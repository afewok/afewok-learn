package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
 * 1584. 连接所有点的最小费用
 */

func Test_leetcode_1584(t *testing.T) {
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
}

func minCostConnectPoints(points [][]int) int {
	defer timeCost()()
	var (
		ans     int = 0
		vis         = map[int]bool{}
		minCost     = make([]int, len(points))
		abs     func(int) int
		min     func(int, int) int
	)
	abs = func(x int) int {
		if x >= 0 {
			return x
		}
		return x * -1
	}
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(minCost); i++ {
		minCost[i] = math.MaxInt64
	}
	minCost[0] = 0
	for i := 0; i < len(points); i++ {
		u := -1
		pMin := math.MaxInt64
		for j := 0; j < len(points); j++ {
			if !vis[j] && minCost[j] < pMin {
				pMin = minCost[j]
				u = j
			}
		}
		ans += pMin
		vis[u] = true
		for j := 0; j < len(points); j++ {
			if !vis[j] {
				cost := abs(points[j][0]-points[u][0]) + abs(points[j][1]-points[u][1])
				minCost[j] = min(minCost[j], cost)
			}
		}
	}
	return ans
}
