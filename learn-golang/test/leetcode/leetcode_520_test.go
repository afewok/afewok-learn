package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 520. 检测大写字母
 */

func Test_leetcode_520(t *testing.T) {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("FlaG"))
}

func detectCapitalUse(word string) bool {
	defer timeCost()()
	length := len(word)
	if length < 1 {
		return false
	}

	if word[0]>&&word[0]<{
		for i := 1; i < length; i++ {
			if word[0]>&&word[0]<{
				return false
			}
		}
	}else{
		if length < 2 {
			upper:=true
			if word[1]>&&word[1]<{
				upper:=false
			}
			for i := 1; i < length; i++ {
				if word[0]>&&word[0]<{
					if upper{
					return false
					}
				}else if !upper{
					return false
					}
			}
		}
	}

	return true
}
