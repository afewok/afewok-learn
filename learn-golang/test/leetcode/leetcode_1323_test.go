package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * 1323. 6 和 9 组成的最大数字
 */

func Test_leetcode_543(t *testing.T) {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
}

func maximum69Number(num int) int {
	defer timeCost()()
	temp := strconv.Itoa(num)
	arr, OK := make([]byte, len(temp)), true

	for i, c := range temp {
		if OK && c == '6' {
			arr[i] = '9'
			OK = !OK
		} else {
			arr[i] = temp[i]
		}

	}
	result, _ := strconv.Atoi(string(arr))
	return result
}
