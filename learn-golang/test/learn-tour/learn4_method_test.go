package learntour

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	// "golang.org/x/tour/pic"
	// "golang.org/x/tour/reader"
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

//接口
func Test_interfaces(t *testing.T) {
	println("接口类型 是由一组方法签名定义的集合。接口不允许指定参数类型是指针，而入参传了值")
	println("接口类型的变量可以保存任何实现了这些方法的值")

	var a abser
	f := MyFloat5(-math.Sqrt2)
	v := Vertex5{3, 4}

	a = f //a MyFloat 实现了abser
	fmt.Println(a.abs5())

	a = &v //a *Vertex 实现了abser
	fmt.Println(a.abs5())

	a = &f //a MyFloat 实现了abser
	fmt.Println(a.abs5())

	// a = v //v 是一个Vertex(不是 *Vertex)，所以没有实现abser

}

type abser interface {
	abs5() float64
}

type MyFloat5 float64

func (f MyFloat5) abs5() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex5 struct {
	X, Y float64
}

func (v *Vertex5) abs5() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//接口与隐式实现
func Test_interfaces_are_satisfied_implicitly(t *testing.T) {
	println("类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。")
	println("隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。")
	println("因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。")

	var i I = &T{"hello"}
	i.M()

}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

//接口值
func Test_interface_values(t *testing.T) {
	println("接口也是值。它们可以像其它值一样传递。")
	println("接口值可以用作函数的参数或返回值。在内部，接口值可以看做包含值和具体类型的元组")
	println("接口值保存了一个具体底层类型的具体值。接口值调用方法时会执行其底层类型的同名方法。")
	var i I = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

//底层值为 nil 的接口值
func Test_interface_values_with_nil(t *testing.T) {
	println("即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。")
	println("在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它")
	println("nil 在 Go语言中只能被赋值给指针和接口")
	println("Golang中的interface类型包含两部分信息——值信息和类型信息，只有interface的值合并类型都为nil时interface才为nil")
	var i I = errFunc()
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func errFunc() *T {
	return nil
}

//nil 接口值
func Test_nil_interface_values(t *testing.T) {
	println("nil 接口值既不保存值也不保存具体类型。")
	println("为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。")
	var i I
	describe(i)
	// i.M() //为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
}

//空接口
func Test_empty_interface(t *testing.T) {
	println("指定了零个方法的接口值被称为 *空接口：*  interface{}")
	println("空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）")
	println("空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。")
	var i interface{}
	describe1(i)

	i = 42
	describe1(i)

	i = "hello"
	describe1(i)
}

func describe1(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}

//类型断言
func Test_type_assertions(t *testing.T) {
	println("类型断言 提供了访问接口值底层具体值的方式  t := i.(T) 若t未保存T类型的值，该语句会发生 恐慌")
	println("")
	println("为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。 t, ok := i.(T)   ok为false时，t为初始化值")
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) //报错（panic）
}

//类型选择
func Test_type_switches(t *testing.T) {
	println("类型与值都可以是属性")
	println("类型选择 是一种按顺序从几个类型断言中选择分支的结构。")
	println("类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。")
	println("类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type。")
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don`t know about type %T!\n", v)
	}
}

//Stringer
func Test_stringer(t *testing.T) {
	println("Stringer 是一个可以用字符串描述自己的类型，类似toString 方法重写")
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z.String())
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

//练习：Stringer
func Test_exercise_stringer(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type IPAddr [4]byte

func (arr IPAddr) String() string {
	// return fmt.Sprintf("%s", strings.Join(arr, "."))

	// fmt.Printf()
	var temp string = ""
	for _, v := range arr {
		if temp == "" {
			temp = fmt.Sprintf("%v", v)
		} else {
			temp = temp + fmt.Sprintf(".%v", v)
		}
	}
	return temp
}

//错误(重写Error方法，类似返回error重写输出)
func Test_errors(t *testing.T) {
	println("Go 程序使用 error 值来表示错误状态")
	println("通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。")
	println("error 为 nil 时表示成功；非 nil 的 error 表示失败。")

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn`t convert number: %v\n", err)
	} else {
		fmt.Println("Converted integer:", i)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn`t work"}
}

//练习：错误
func Test_exercise_errors(t *testing.T) {
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt1(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	println("在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么呢？")
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt1(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

//Reader
func Test_reader(t *testing.T) {
	println("io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。")
	println("io.Reader 接口有一个 Read 方法： func (T) Read(b []byte) (n int, err error)")
	println("Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。")

	r := strings.NewReader("Hello World")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

//练习：Reader
// func Test_exercise_reader(t *testing.T) {
// 	println("实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。")
// 	reader.Validate(MyReader{})
// }

func (m MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

type MyReader struct{}

//练习：rot13Reader
func Test_exercise_rot_reader(t *testing.T) {
	println("有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。")
	println("例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。")
	println("编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。")

	s := strings.NewReader("Lbh penpxrq gur pdqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (int, error) {
	n, e := rot.r.Read(p)
	for i := 0; i < n; i++ {
		b := p[i]
		switch {
		case (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm'):
			p[i] += 13
		case (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z'):
			p[i] -= 13
		}
	}
	return n, e
}

//图像
func Test_images(t *testing.T) {
	println("color.Color 和 color.Model 类型也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。这些接口和类型由 image/color 包定义。")
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

//练习：图像
// func Test_exercise_images(t *testing.T) {
// 	m := Image{}
// 	pic.ShowImage(m)
// }

type Image struct{} //新建一个Image结构体

func (i Image) ColorModel() color.Model {
	//实现Image包中颜色模式的方法
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	//实现Image包中生成图片边界的方法
	return image.Rect(0, 0, 200, 200)
}

func (i Image) At(x, y int) color.Color {
	//实现Image包中生成图像某个点的方法
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}
