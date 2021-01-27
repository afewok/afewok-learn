package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 637. 二叉树的层平均值
 */

func Test_leetcode_637(t *testing.T) {
	fmt.Println(averageOfLevels(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
}

func averageOfLevels(root *TreeNode) []float64 {
	defer timeCost()()
	arrays, average := make([][]int, 0), make([]float64, 0)
	getAverageOfLevels(root, &arrays, 0)
	for _, array := range arrays {
		average = append(average, float64(array[0])/float64(array[1]))
	}
	return average
}

func getAverageOfLevels(node *TreeNode, arr *[][]int, tier int) {
	if node == nil {
		return
	}

	if len(*arr) <= tier {
		*arr = append(*arr, make([]int, 2))
	}
	arrays := *arr
	arrays[tier][0] += node.Val
	arrays[tier][1]++
	getAverageOfLevels(node.Left, arr, tier+1)
	getAverageOfLevels(node.Right, arr, tier+1)
}
