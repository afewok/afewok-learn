package learntour

import (
	"fmt"
	"math"
	"testing"
)

//方法
func Test_methods(t *testing.T) {
	println("Go 没有类。不过你可以为结构体类型定义方法。方法就是一类带特殊的 接收者 参数的函数")
	println("方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间")
	v := VertexA{3, 4}
	fmt.Println(v.abs1())
}

type VertexA struct {
	X, Y float64
}

func (v VertexA) abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//方法即函数
func Test_Methods_funcs(t *testing.T) {
	println("方法只是个带接收者参数的函数")
	v := VertexA{3, 4}
	fmt.Println(abs2(v))
}

func abs2(v VertexA) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//方法（续）
func Test_methods_continued(t *testing.T) {
	println("也可以为非结构体类型声明方法")
	println("只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法")
	println("接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs3())
}

type MyFloat float64

func (f MyFloat) abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//指针接收者
func Test_methods_pointers(t *testing.T) {
	println("指针接受者会修改原始值，对原始值产生副作用")
	v := VertexA{3, 4}
	v.scale1(10)
	fmt.Println(v.abs1())
}

func (v *VertexA) scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//指针与函数
func Test_methods_pointers_explained(t *testing.T) {
	v := VertexA{3, 4}
	scale2(&v, 10)
	fmt.Println(abs2(v))
}

func scale2(v *VertexA, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//方法与指针重定向
func Test_indirection(t *testing.T) {
	println("带指针参数的函数必须接受一个指针,而以指针为接收者的方法被调用时，接收者既能为值又能为指针")
	println("原因，为方便 Go会带指针接受者检测并加上& ----> v.scale1()、(&v).scale1()")
	v := VertexA{3, 4}
	v.scale1(2)
	scale2(&v, 10)

	p := &VertexA{4, 3}
	p.scale1(3)
	scale2(p, 8)
	fmt.Println(v, p)
}

//方法与指针重定向（续）
func Test_indirection_values(t *testing.T) {
	println("函数入参不是指针，那么传参也不能是指针，但方法接收者不受限制，传参既可以指针也可以不是")
	println("原因，为方便 Go会带指针接受者检测并加上& ----> v.scale1()、(&v).scale1()")
	println()
	v := VertexA{3, 4}
	fmt.Println(v.abs3(), abs3(v))

	p := &VertexA{4, 3}
	fmt.Println(p.abs3())
	// fmt.Println(abs3(p))//编译报错
}

func (v VertexA) abs3() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func abs3(v VertexA) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//选择值或指针作为接收者
func Test_methods_with_pointer_receivers(t *testing.T) {
	println("使用指针接收者的原因有二：")
	println("首先，方法能够修改其接收者指向的值。")
	println("其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。")
	println("通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。")

	v := &VertexA{3, 4}
	fmt.Printf("Before scaling:%+v,Abs:%v\n", v, v.abs4())
	v.scale4(5)
	fmt.Printf("After scaling:%+v,Abs:%v\n", v, v.abs4())
}

func (v *VertexA) scale4(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *VertexA) abs4() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
