package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5621. 无法吃午餐的学生数量
 */

func Test_leetcode_5621(t *testing.T) {
	// fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	// fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
	fmt.Println(countStudents([]int{0, 0, 0, 1, 1, 1, 1, 0, 0, 0}, []int{1, 0, 1, 0, 0, 1, 1, 0, 0, 0}))
}
func countStudents(students []int, sandwiches []int) int {
	defer timeCost()()
	temp, num, student, sandwiche, lenStudent, lenSandwiche := 0, 0, 0, 0, len(students), len(sandwiches)
	for sandwiche < lenSandwiche {
		if students[student] == sandwiches[sandwiche] {
			students[student] = -1
			sandwiches[sandwiche] = -1
			sandwiche++
			temp++
		}

		student++
		if student == lenStudent {
			student = 0
			if num == temp {
				return lenStudent - temp
			}
			num = temp
		}
	}
	return lenStudent - temp
}
