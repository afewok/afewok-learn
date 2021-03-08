package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 1683. 无效的推文
 */

func Test_leetcode_1683(t *testing.T) {
	fmt.Println("", "select tweet_id  from Tweets where LENGTH(content)>15")
	fmt.Println("", "select tweet_id  from Tweets where char_length(content)>15")
}
