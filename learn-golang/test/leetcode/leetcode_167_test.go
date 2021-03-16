package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 167. 两数之和 II - 输入有序数组
 */

func Test_leetcode_167(t *testing.T) {
	fmt.Println(twoSum167([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum167([]int{2, 3, 4}, 6))
	fmt.Println(twoSum167([]int{-1, 0}, -1))
}

//双指针
func twoSum167(numbers []int, target int) []int {
	defer timeCost()()
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

//二分
func twoSum167_1(numbers []int, target int) []int {
	defer timeCost()()
	for i, num := range numbers {
		left, right := 0, len(numbers)
		for left < right {
			mid := left + (left+right+1)/2
			if num+numbers[mid] == target {
				return []int{i + 1, mid + 1}
			} else if num+numbers[mid] > target {
				right = mid - 1
			} else if num+numbers[mid] < target {
				left = mid + 1
			}
		}
	}
	return []int{}
}

func twoSum167_2(numbers []int, target int) []int {
	m := map[int]int{}
	for i, v := range numbers {
		find := target - v
		if _, e := m[find]; e {
			return []int{m[find] + 1, i + 1}
		}
		m[v] = i
	}
	return []int{}
}
