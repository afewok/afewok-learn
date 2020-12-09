package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 62. 不同路径
 *
 * 一个机器人位于一个 m x n 网格的左上角 。 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角。问总共有多少条不同的路径？
 *
 * 示例 1:输入: m = 3, n = 2 输出: 3 解释:从左上角开始，总共有 3 条路径可以到达右下角。1. 向右 -> 向右 -> 向下2. 向右
 * -> 向下 -> 向右3. 向下 -> 向右 -> 向右
 *
 * 示例 2:输入: m = 7, n = 3 输出: 28
 *
 * 示例 3：输入：m = 3, n = 3 输出：6
 *
 * 提示：1 <= m, n <= 100 题目数据保证答案小于等于 2 * 10 ^ 9
 *
 * 思路：二叉树遍历，控制变量法
 */
func Test_leetcode_062(t *testing.T) {
	fmt.Println(uniquePaths1(10, 10))
	fmt.Println(uniquePaths2(3, 6))
}

func uniquePaths1(m int, n int) int {
	if m == 1 {
		return 1
	}
	var arrays []int = make([]int, n)
	for i, v := range arrays {
		arrays[i] = v + 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arrays[j] = arrays[j-1] + arrays[j]
		}
	}
	return arrays[n-1]
}

func uniquePaths2(m int, n int) int {
	return uniquePaths(1, 1, m, n)
}

func uniquePaths(y, x, m, n int) int {
	if x > n || y > m {
		return 0
	} else if x == n && y == m {
		return 1
	}
	return uniquePaths(y+1, x, m, n) + uniquePaths(y, x+1, m, n)
}
