package leetcode

import (
	"fmt"
	"testing"
)

func Test_leetcode_860(t *testing.T) {
	fmt.Println(lemonadeChange2([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange2([]int{5, 5, 10}))
	fmt.Println(lemonadeChange2([]int{10, 10}))
	fmt.Println(lemonadeChange2([]int{5, 5, 10, 10, 20}))
}

func lemonadeChange1(bills []int) bool {
	m5, m10, length := 0, 0, len(bills)
	for i := 0; i < length; i++ {
		if bills[i] == 5 {
			m5++
		} else if bills[i] == 10 {
			if m5 > 0 {
				m5--
				m10++
			} else {
				return false
			}
		} else {
			if m5 > 0 && m10 > 0 {
				m5--
				m10--
			} else if m5 > 2 {
				m5 = m5 - 3
			} else {
				return false
			}
		}
	}
	return true
}

func lemonadeChange2(bills []int) bool {
	m5, m10 := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			m5++
		case 10:
			if m5 > 0 {
				m5--
				m10++
			} else {
				return false
			}
		case 20:
			if m5 > 0 && m10 > 0 {
				m5--
				m10--
			} else if m5 > 2 {
				m5 = m5 - 3
			} else {
				return false
			}
		}
	}
	return true
}

func lemonadeChange3(bills []int) bool {
	m5, m10, length := 0, 0, len(bills)
	for i := 0; i < length; i++ {
		if bills[i] == 5 {
			m5++
		} else if bills[i] == 10 {
			m5--
			m10++
		} else if m10 > 0 {
			m5--
			m10--
		} else {
			m5 = m5 - 3
		}

		if m5 < 0 || m10 < 0 {
			return false
		}
	}
	return true
}

func lemonadeChange4(bills []int) bool {
	array := make([]int, 5)
	for _, bill := range bills {
		array[bill/5]++
		array[bill/10]--
		array[bill/20]--
		if array[1] < 0 || array[1]+2*array[2] < 0 {
			return false
		}
	}
	return true
}
