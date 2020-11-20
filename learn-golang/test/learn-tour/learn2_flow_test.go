package learntour

import (
	"fmt"
	"math"
	"testing"
)

//for
func Test_for(t *testing.T) {
	fmt.Println("只有一种循环结构 for，没有小括号，而大括号是必须")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//for（续）
func Test_for_continued(t *testing.T) {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

//for 是 Go 中的 “while”
func Test_for_is_gos_while(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//无限循环
func Test_forever(t *testing.T) {
	count := 0
	for {
		count++
		if count > 100 {
			break
		}
	}
	fmt.Println(count)
}

//if
func Test_if(t *testing.T) {
	fmt.Println(sqrt(2), sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//if 的简短语句
func Test_if_with_a_short_statment(t *testing.T) {
	println("同 for 一样， if 语句可以在条件表达式前执行一个简单的语句")
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
