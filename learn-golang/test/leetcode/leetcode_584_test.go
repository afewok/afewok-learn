package leetcode

import (
	"fmt"
	"testing"
)

/**
 * 584. 寻找用户推荐人
 */

func Test_leetcode_584(t *testing.T) {
	fmt.Println("", "select name from customer where referee_id is null or referee_id!=2")
	fmt.Println("ifnull", "select name from customer where ifnull(referee_id,0)!=2")
}
