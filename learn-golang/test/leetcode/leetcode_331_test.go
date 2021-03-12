package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 331. 验证二叉树的前序序列化
 */

func Test_leetcode_331(t *testing.T) {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization("1,#"))
	fmt.Println(isValidSerialization("9,#,#,1"))
}

func isValidSerialization1(preorder string) bool {
	defer timeCost()()
	n := len(preorder)
	stk := []int{1}
	for i := 0; i < n; {
		if len(stk) == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			stk[len(stk)-1]--
			if stk[len(stk)-1] == 0 {
				stk = stk[:len(stk)-1]
			}
			i++
		} else {
			// 读一个数字
			for i < n && preorder[i] != ',' {
				i++
			}
			stk[len(stk)-1]--
			if stk[len(stk)-1] == 0 {
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, 2)
		}
	}
	return len(stk) == 0
}
func isValidSerialization(preorder string) bool {
	defer timeCost()()
	n := len(preorder)
	slots := 1
	for i := 0; i < n; {
		if slots == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			slots--
			i++
		} else {
			// 读一个数字
			for i < n && preorder[i] != ',' {
				i++
			}
			slots++ // slots = slots - 1 + 2
		}
	}
	return slots == 0
}
