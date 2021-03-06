package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 135. 分发糖果
 */

func Test_leetcode_135(t *testing.T) {
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
	fmt.Println(candy([]int{1, 3, 2, 2, 1}))
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))
	fmt.Println(candy([]int{29, 51, 87, 87, 72, 12}))
	fmt.Println(candy([]int{1, 6, 10, 8, 7, 3, 2}))
	fmt.Println(candy([]int{1, 3, 4, 5, 2}))
}

func candy(ratings []int) int {
	defer timeCost()()
	length, award := len(ratings), 0
	arr := make([]int, length)

	for i := 1; i < length; i++ {
		if ratings[i-1] < ratings[i] && arr[i-1] >= arr[i] {
			arr[i] = arr[i-1] + 1
		}
	}
	award += arr[length-1]
	for i := length - 2; i >= 0; i-- {
		if ratings[i+1] < ratings[i] && arr[i+1] >= arr[i] {
			arr[i] = arr[i+1] + 1
		}
		award += arr[i]
	}
	return length + award
}

func candy1(ratings []int) int {
	defer timeCost()()
	length := len(ratings)
	arr := make([]int, length)

	for i := 1; i < length; i++ {
		if ratings[i-1] > ratings[i] {
			j := i + 1
			for j < length {
				if ratings[j-1] <= ratings[j] {
					break
				}
				j++
			}
			count := 1
			for k := j - 2; k >= i; k-- {
				arr[k] = count
				count++
			}
			if i-1 == 0 || ratings[i-1] == ratings[i-2] {
				arr[i-1] = count
			} else if ratings[i-1] > ratings[i-2] && count > arr[i-1] {
				arr[i-1] = count
			}
			i = j - 1
		} else if ratings[i-1] < ratings[i] {
			arr[i] = arr[i-1] + 1
		}
	}
	for _, v := range arr {
		length += v
	}
	return length
}

func candy2(ratings []int) int {
	defer timeCost()()
	length, award, sub := len(ratings), 0, -1
	arr := make([]int, length)
	for i := 1; i < length; i++ {
		if ratings[i-1] < ratings[i] {
			arr[i] = arr[i-1] + 1
			award += arr[i]
			sub = -1
		} else if ratings[i-1] > ratings[i] {
			if sub == -1 {
				sub = i - 1
			} else if sub+1 < i {
				for j := i - 1; j > sub; j-- {
					arr[j]++
					award++
				}
			}
			if arr[sub] == arr[sub+1] {
				arr[sub]++
				award++
			}
		} else {
			sub = -1
		}
	}
	return length + award
}
