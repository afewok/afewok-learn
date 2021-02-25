package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 426. 将二叉搜索树转化为排序的双向链表
 */

func Test_leetcode_426(t *testing.T) {
	fmt.Println(treeToDoublyList(&TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}))
}

func treeToDoublyList(root *TreeNode) *Node {
	defer timeCost()()

}
