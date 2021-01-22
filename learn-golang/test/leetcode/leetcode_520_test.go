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
	fmt.Println(detectCapitalUse("ffffffffffffffffffffF"))
}

func detectCapitalUse(word string) bool {
	defer timeCost()()
	length := len(word)
	if length < 1 {
		return false
	}

	if word[0] >= 97 && word[0] <= 122 {
		for i := 1; i < length; i++ {
			if word[i] < 97 || word[i] > 122 {
				return false
			}
		}
	} else {
		if length > 2 {
			upper := true
			if word[1] >= 97 && word[1] <= 122 {
				upper = false
			}
			for i := 2; i < length; i++ {
				if word[i] >= 97 && word[i] <= 122 {
					if upper {
						return false
					}
				} else if !upper {
					return false
				}
			}
		}
	}
	return true
}
