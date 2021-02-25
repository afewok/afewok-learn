package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 230. 二叉搜索树中第K小的元素
 */

func Test_leetcode_230(t *testing.T) {
	fmt.Println(kthSmallest(&TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}, 1))
	fmt.Println(kthSmallest(&TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}, 4))
}

var result int

func kthSmallest(root *TreeNode, k int) int {
	defer timeCost()()
	getKthSmallest(root, k, 0)
	return result
}

func getKthSmallest(root *TreeNode, k, count int) int {
	if root.Left != nil {
		count = getKthSmallest(root.Left, k, count)
		if k == count {
			return count
		}
	}
	count++
	if k == count {
		result = root.Val
		return count
	}

	if root.Right != nil {
		count = getKthSmallest(root.Right, k, count)
		if k == count {
			return count
		}
	}

	return count
}

func kthSmallestII(root *TreeNode, nodenum_root *TreeNode, k int) int {
	position := nodenum_root.Val
	if root.Right != nil {
		position -= nodenum_root.Right.Val
	}
	if position == k {
		return root.Val
	} else if position > k {
		return kthSmallestII(root.Left, nodenum_root.Left, k)
	} else {
		return kthSmallestII(root.Right, nodenum_root.Right, k-position)
	}
}
