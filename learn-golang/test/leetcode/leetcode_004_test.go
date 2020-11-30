package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 4. 寻找两个正序数组的中位数
 *
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。 进阶：你能设计一个时间复杂度为
 * O(log (m+n)) 的算法解决此问题吗？
 *
 * 示例 1：输入：nums1 = [1,3], nums2 = [2] 输出：2.00000 解释：合并数组 = [1,2,3] ，中位数 2
 *
 * 示例 2：输入：nums1 = [1,2], nums2 = [3,4] 输出：2.50000 解释：合并数组 = [1,2,3,4] ，中位数 (2 +
 * 3) / 2 = 2.5
 *
 * 示例 3：输入：nums1 = [0,0], nums2 = [0,0] 输出：0.00000
 *
 * 示例 4：输入：nums1 = [], nums2 = [1] 输出：1.00000
 *
 * 示例 5：输入：nums1 = [2], nums2 = [] 输出：2.00000
 *
 * 提示：nums1.length == m nums2.length == n 0 <= m <= 1000 0 <= n <= 1000 1 <= m +
 * n <= 2000 -106 <= nums1[i], nums2[i] <= 106
 *
 * 思路：合并排序取值、下标记录、二分查找、划分数组
 */
func Test_leetcode_004(t *testing.T) {
	var num1, num2 []int = []int{1, 2}, []int{3, 4}
	fmt.Println(findMedianSortedArrays1(num1, num2))
	fmt.Println(findMedianSortedArrays2(num1, num2))
	fmt.Println(findMedianSortedArrays3(num1, num2))
	fmt.Println(findMedianSortedArrays4(num1, num2))

}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	var pair int = length / 2
	var list []int = append(nums1, nums2...)
	sort.Sort(sort.IntSlice(list))
	if pair*2 == length {
		return float64(list[pair-1]+list[pair]) / 2.0
	}
	return float64(list[pair])
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	nums1Length, nums2Length := len(nums1), len(nums2)
	totalLength := nums1Length + nums2Length
	pair, last, current := totalLength/2, 0, 0
	for i, n1, n2 := 0, 0, 0; i < totalLength; i++ {
		last = current
		if nums1Length > n1 && nums2Length > n2 {
			if nums1[n1] <= nums2[n2] {
				current = nums1[n1]
				n1++
			} else {
				current = nums2[n2]
				n2++
			}
		} else if nums1Length > n1 {
			current = nums1[n1]
			n1++
		} else {
			current = nums2[n2]
			n2++
		}

		if pair == i {
			if pair*2 == totalLength {
				return float64(last+current) / 2.0
			}
			return float64(current)
		}
	}
	panic("Illegal nums1 or nums2")
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {

	return 0
}

func findMedianSortedArrays4(nums1 []int, nums2 []int) float64 {

	return 0
}
