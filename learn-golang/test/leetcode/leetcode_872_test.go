package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 872. 叶子相似的树
 *
 * 思路：
 */
func Test_leetcode_872(t *testing.T) {
	root1 := &TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}}, &TreeNode{1, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}}}
	root2 := &TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}, &TreeNode{1, &TreeNode{4, nil, nil}, &TreeNode{2, &TreeNode{9, nil, nil}, &TreeNode{8, nil, nil}}}}
	fmt.Println(leafSimilar(root1, root2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	defer timeCost()()
	arr1, arr2 := make([]int, 0), make([]int, 0)
	getLeaf(root1, &arr1)
	getLeaf(root2, &arr2)
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

func getLeaf(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	} else if node.Left == nil && node.Right == nil {
		*arr = append(*arr, node.Val)
		return
	}
	getLeaf(node.Left, arr)
	getLeaf(node.Right, arr)
}
