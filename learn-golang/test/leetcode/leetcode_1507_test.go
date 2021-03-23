package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1507. 转变日期格式
 */

func Test_leetcode_1507(t *testing.T) {
	fmt.Println(reformatDate("20th Oct 2052"))
	fmt.Println(reformatDate("6th Jun 1933"))
	fmt.Println(reformatDate("26th May 1960"))
}

func reformatDate(date string) string {
	defer timeCost()()
	length := len(date)
	i, arr := 0, make([]string, 2)
	for ; i < length; i++ {
		if date[i] < '0' || date[i] > '9' {
			arr[0] = date[:i]
			break
		}
	}
	for ; i < length; i++ {
		if date[i] == ' ' {
			break
		}
	}
	j := i + 1
	for ; j < length; j++ {
		if date[j] == ' ' {
			arr[1] = date[i+1 : j]
			break
		}
	}
	if len(arr[0]) == 1 {
		return date[j+1:] + "-" + getMonth1507(arr[1]) + "-0" + arr[0]
	}
	return date[j+1:] + "-" + getMonth1507(arr[1]) + "-" + arr[0]
}

func getMonth1507(str string) string {
	switch str {
	case "Jan":
		return "01"
	case "Feb":
		return "02"
	case "Mar":
		return "03"
	case "Apr":
		return "04"
	case "May":
		return "05"
	case "Jun":
		return "06"
	case "Jul":
		return "07"
	case "Aug":
		return "08"
	case "Sep":
		return "09"
	case "Oct":
		return "10"
	case "Nov":
		return "11"
	case "Dec":
		return "12"
	}
	return ""
}
