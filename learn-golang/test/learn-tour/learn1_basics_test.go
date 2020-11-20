package learntour

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"testing"
	"time"
)

func Test_hello(t *testing.T) {
	fmt.Println("Hello, 世界. 发布于 2009-11-10 23:00:00 UTC")
}

//当前时间
func Test_sandbox(t *testing.T) {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
}

//包
func Test_package(t *testing.T) {
	println("随机数种子")
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}

//导入
func Test_import(t *testing.T) {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

//导出名
func Test_exported(t *testing.T) {
	println("大写字母开头 ，表示public可调用")
	fmt.Println(math.Pi)
}

//函数
func Test_functions(t *testing.T) {
	fmt.Println(add1(42, 13))
}

func add1(x int, y int) int {
	return x + y
}

//函数（续）
func Test_functions_continued(t *testing.T) {
	fmt.Println(add2(42, 13))
}

func add2(x, y int) int {
	return x + y
}

//多值返回
func Test_multiple_results(t *testing.T) {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}

//命名返回值
func Test_named_results(t *testing.T) {
	fmt.Println(split(17))
}

func split(sum int) (x, y int) {
	println("没有参数的 return 语句返回已命名的返回值。也就是 直接 返回")
	x = sum * 4 / 9
	y = sum - x
	return
}

//变量
func Test_variables(t *testing.T) {
	var i int
	fmt.Println(i, c, python, java)
}

var c, python, java bool

//变量的初始化
func Test_variables_with_initializers(t *testing.T) {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

var i, j int = 1, 2

//短变量声明
func Test_short_variable_declarations(t *testing.T) {
	println("函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。")
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

//基础类型
func Test_basic_types(t *testing.T) {
	println("变量可以分组")
	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//零值
func Test_zero(t *testing.T) {
	println("字符串默认值 \"\"")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q \n", i, f, b, s)
}

//类型转化
func Test_type_conversions(t *testing.T) {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}

//类型推导
func Test_type_inference(t *testing.T) {
	i := 42           //int
	f := 3.142        //float64
	g := 0.867 + 0.5i //complex128
	fmt.Println(i, f, g)
}

//常量
func Test_constants(t *testing.T) {
	println("常量不能用 := 语法声明")
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const Pi = 3.14

//数值常量
func Test_numeric_constants(t *testing.T) {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

const (
	//将 1 左移100位来创建一个非常大的数字
	//即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	//再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
