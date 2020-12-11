package leetcode

import "testing"

/**
* 649. Dota2 参议院
 */
/**
 * 376. 摆动序列
 *
 * 思路：贪婪循环遍历（先记录，遇到在毙掉）
 */
func Test_leetcode_649(t *testing.T) {
	predictPartyVictory("DRRDRDRDRDDRDRDR")
}

func predictPartyVictory(senate string) string {
	R, D := true, true
	bytes := []byte(senate)
	flag, length := 0, len(bytes)
	for R && D {
		R, D = false, false
		for i := 0; i < length; i++ {
			if bytes[i] == 'R' {
				if flag < 0 {
					bytes[i] = 0
				} else {
					R = true
				}
				flag++
			} else if bytes[i] == 'D' {
				if flag > 0 {
					bytes[i] = 0
				} else {
					D = true
				}
				flag--
			}

		}
	}

	if R {
		return "Radiant"
	}
	return "Dire"
}
