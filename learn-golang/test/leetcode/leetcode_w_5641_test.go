package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 5641. 卡车上的最大单元数
 */

func Test_leetcode_5641(t *testing.T) {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	defer timeCost()()
	length := len(boxTypes)
	arr := make([][]int, length)
	for i := 0; i < length; i++ {
		arr[i] = []int{boxTypes[i][1], boxTypes[i][0]}
		for j := i; j > 0; j-- {
			if arr[j][0] > arr[j-1][0] {
				arr[j][0], arr[j-1][0] = arr[j-1][0], arr[j][0]
				arr[j][1], arr[j-1][1] = arr[j-1][1], arr[j][1]
			}
		}
	}

	result := 0
	for i := 0; i < length && truckSize > 0; i++ {
		if arr[i][1] <= truckSize {
			result = result + arr[i][1]*arr[i][0]
			truckSize = truckSize - arr[i][1]
		} else {
			result = result + truckSize*arr[i][0]
			truckSize = 0
		}
	}
	return result
}
