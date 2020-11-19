package learntour

import (
	"fmt"
	"testing"
)

func Test_for(t *testing.T) {
	fmt.Println("只有一种循环结构 for，没有小括号，而大括号是必须")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
