package learntour

import (
	"fmt"
	"math"
	"runtime"
	"testing"
	"time"
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
	} else {
		return lim
	}
}

//if 和 else
func Test_if_and_else(t *testing.T) {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

//练习：循环与函数，用牛顿法实现平方根函数
func Test_exercise_loops_and_functions(t *testing.T) {
	println("计算机通常使用循环来计算 x 的平方根。从某个猜测的值 z 开始，我们可以根据 z² 与 x 的近似度来调整 z，产生一个更好的猜测：")
	println("z -= (z*z - x) / (2*z)")
	println("重复调整的过程，猜测的结果会越来越精确，得到的答案也会尽可能接近实际的平方根。")
	fmt.Println(Sqrt(2))
}

func Sqrt(x float64) float64 {
	z := float64(1)    // 定义一个初始值并对它初始化
	temp := float64(0) //临时变量，作为记录z 上次的值
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(z-temp) < 0.000000000000001 {
			return z
		} else {
			temp = z
		}
	}
}

//switch
func Test_switch(t *testing.T) {
	println("switch 的 case 无需为常量，且取值不必为整数。")
	println("Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("windows.")
		fallthrough
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s. \n", os)
	}
}

//switch 的求值顺序
func Test_switch_evaluation_order(t *testing.T) {
	println("switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。")
	fmt.Println("When`s Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//没有条件的 switch
func Test_switch_with_no_condition(t *testing.T) {
	println("没有条件的 switch 同 switch true 一样。这种形式能将一长串 if-then-else 写得更加清晰。")
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Good morning!")
	case time.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

//defer
func Test_defer(t *testing.T) {
	defer fmt.Println("world")
	fmt.Print("hello ")
}

//defer 栈
func Test_defer_multi(t *testing.T) {
	print("推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。")
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
