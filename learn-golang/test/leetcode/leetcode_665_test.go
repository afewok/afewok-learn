package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 665. 非递减数列
 */

func Test_leetcode_665(t *testing.T) {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
}

func checkPossibility(nums []int) bool {
	defer timeCost()()

}
