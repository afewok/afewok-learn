package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 219. 存在重复元素 II
 */

func Test_leetcode_219(t *testing.T) {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	defer timeCost()()

}
