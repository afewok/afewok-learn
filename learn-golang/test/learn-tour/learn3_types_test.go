package learntour

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

//指针
func Test_pointers(t *testing.T) {
	println("指针保存了值的内存地址。与 C 不同，Go 没有指针运算。")
	println("类型 *T 是指向 T 类型值的指针。其零值为 nil。")
	println("& 操作符会生成一个指向其操作数的指针。")
	println("* 操作符表示指针指向的底层值。这也就是通常所说的“间接引用”或“重定向”。")

	i, j := 42, 2701
	p := &i         //指向i的指针
	fmt.Println(*p) //通过指针读取 i 的值
	*p = 21         //通过指针设置 i 的值
	fmt.Println(*p) //查看 i 的值

	p = &j         //指向 j
	*p = *p / 37   //通过指针对 j 进行除法运算
	fmt.Println(j) //查看 j 的值

	fmt.Println(p) //查看指针 p （输出内存地址）
}

//结构体
func Test_structs(t *testing.T) {
	println("一个结构体（struct）就是一组字段（field）")
	fmt.Println(Vertex{1, 2})
}

type Vertex struct {
	X int
	Y int
}

//结构体字段
func Test_struct_fields(t *testing.T) {
	println("结构体字段使用点号来访问")

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

//结构体指针
func Test_struct_pointers(t *testing.T) {
	println("可以通过 (*p).X 来访问其字段 X,隐式简写 p.X")
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

//结构体文法
func Test_struct_literals(t *testing.T) {
	println("使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）")
	println("特殊的前缀 & 返回一个指向结构体的指针。结构体不能直接返回内存地址")
	fmt.Println(v1, v2, v3, p, *p, &v1)
}

type Vertex1 struct {
	X, Y int
}

var (
	v1 = Vertex1{1, 2}  //创建一个 Vertex1 类型的结构体
	v2 = Vertex1{X: 1}  //Y:0 被隐式地赋予
	v3 = Vertex1{}      //X:0 Y:0
	p  = &Vertex1{3, 4} //创建一个 *Vertex1 类型的结构体（指针）
)

//数组
func Test_array(t *testing.T) {
	println("类型 [n]T 表示拥有 n 个 T 类型的值的数组。")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

//切片
func Test_slices(t *testing.T) {
	println("每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。")
	println("类型 []T 表示一个元素类型为 T 的切片,切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：a[low : high] 左闭右开")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[0:4]
	fmt.Println(s)
}

//切片就像数组的引用
func Test_slices_pointers(t *testing.T) {
	names := [4]string{"John", "Paul", "Genrge", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

//切片文法
func Test_slice_literals(t *testing.T) {
	println("匿名结构体和初始值")
	q := []int{2, 3, 4, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)

	vertexs := []Vertex{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(vertexs)
}

//切片的默认行为
func Test_slice_bounds(t *testing.T) {
	println("在进行切片时，你可以利用它的默认行为来忽略上下界。切片下界的默认值为 0，上界则是该切片的长度。")
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[:]
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

//切片的长度与容量
func Test_slice_len_cap(t *testing.T) {
	println("切片拥有 长度 和 容量")
	println("切片的长度就是它所包含的元素个数,切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数")
	println("长度和容量可通过表达式 len(s) 和 cap(s) 来获取,可以通过重新切片来扩展一个切片")
	println("前切会被舍弃，无法恢复")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] //截取切片使其长度为0
	printSlice(s)

	s = s[:4] //拓展其长度
	printSlice(s)

	s = s[2:] //舍弃前两个值
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//nil 切片
func Test_nil_slices(t *testing.T) {
	println("切面的默认值是nil，nil 切片的长度和容量为 0 且没有底层数组")
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

//用 make 创建切片
func Test_making_slices(t *testing.T) {
	println("切片可以用内建函数 make 来创建，这也是你创建动态数组的方式")

	a := make([]int, 5) //len(a)=5
	printSlice(a)

	b := make([]int, 0, 5) //len(b)=0,cap(b)=5
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

//切片的切片
func Test_slices_of_slices(t *testing.T) {
	println("切片可包含任何类型，甚至包括其它的切片")
	//创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	//两个玩家轮流打上 X 或 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

//向切片追加元素
func Test_append(t *testing.T) {
	println("为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数  func append(s []T, vs ...T) []T")
	println("append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾")
	println("append 的结果是一个包含原切片所有元素加上新添加元素的切片")
	println("当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。")
	println("容量足够则不会分配新容量，而是直接覆盖原容量的值")
	var s []int
	s = append(s, 0) //添加一个空切片
	printSlice(s)

	s = append(s, 1) //这个切片会按需增加
	printSlice(s)

	s = append(s, 2, 3, 4) //可以一次性添加多个元素
	printSlice(s)

	s = s[:1] //切片保留1个元素
	printSlice(s)

	d := append(s, 5, 6, 7, 8) //再加元素，此时cap不变
	s = s[0:6]                 //而原切面的空间数据被覆盖了
	printSlice(d)
	printSlice(s)

	d = d[2:] //头部舍弃2个元素
	printSlice(d)
}

//Range
func Test_range(t *testing.T) {
	println("for 循环的 range 形式可遍历切片或映射")
	println("当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。")
	for i, v := range pows {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}

var pows = []int{1, 2, 4, 8, 16, 32, 64, 128}

//range（续）
func Test_range_continued(t *testing.T) {
	println("可以将下标或值赋予 _ 来忽略它")
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) //==2**i
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}

//练习：切片
func Test_exercise_slices(t *testing.T) {
	pic.Show(Pic(func(x, y int) int {
		return (x + y) / 2
	}))

	pic.Show(Pic(func(x, y int) int {
		return x * y
	}))

	pic.Show(Pic(func(x, y int) int {
		return x ^ y
	}))

	pic.Show(Pic(func(x, y int) int {
		return int(float64(x) * math.Log(float64(y)))
	}))

	pic.Show(Pic(func(x, y int) int {
		return x % (y + 1)
	}))
}

func Pic(fn func(x, y int) int) func(dx, dy int) [][]uint8 {

	return func(dx, dy int) [][]uint8 {
		var pic [][]uint8 = make([][]uint8, dy)
		for y := 0; y < dy; y++ {
			pic[y] = make([]uint8, dy)
			for x := 0; x < dx; x++ {
				pic[y][x] = uint8(fn(x, y))
			}

		}
		return pic
	}
}

//映射
func Test_maps(t *testing.T) {
	println("映射的零值为 nil 。nil 映射既没有键，也不能添加键")
	println("make 函数会返回给定类型的映射，并将其初始化备用。")
	m = make(map[string]VertexM)
	m["Bell Labs"] = VertexM{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

type VertexM struct {
	Lat, Long float64
}

var m map[string]VertexM

//映射的文法
func Test_map_literals(t *testing.T) {
	println("映射的文法与结构体相似，不过必须有键名。")
	println("map的分割符是：，结构体的是逗号")
	var m = map[string]VertexM{
		"Bell Labs": VertexM{
			40.68433, -74.39967,
		},
		"Google": VertexM{
			37.42202, -122.08408,
		},
		"CCC": VertexM{1, 2},
	}
	fmt.Println(m)
}

//映射的文法（续）
func Test_map_literals_continued(t *testing.T) {
	println("若顶级类型只是一个类型名，你可以在文法的元素中省略它。")
	var m = map[string]VertexM{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

//修改映射
func Test_mutating_maps(t *testing.T) {
	m := make(map[string]int)
	fmt.Println(m)
	m = map[string]int{}
	fmt.Println(m)

	m["Answer"] = 42                       //在映射 m 中插入或修改元素
	fmt.Println("The value:", m["Answer"]) //获取元素

	m["Answer"] = 48                       //在映射 m 中插入或修改元素
	fmt.Println("The value:", m["Answer"]) //获取元素

	delete(m, "Answer") //删除元素
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] //通过双赋值检测某个键是否存在,不存在则v=该类型的默认值，ok为false
	fmt.Println("The Value:", v, "Present?", ok)
}

//练习：映射
func Test_exercise_maps(t *testing.T) {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	maps := map[string]int{}
	for _, v := range strs {
		maps[v]++
	}
	return maps
}

//函数值
func Test_function_values(t *testing.T) {
	fmt.Println("函数也是值。它们可以像其它值一样传递。")
	fmt.Println("函数值可以用作函数的参数或返回值。")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println(func(x, y int) int {
		return x + y
	}(1, 2))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//函数的闭包
func Test_function_closures(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//练习：斐波纳契闭包（这个数列从第3项开始，每一项都等于前两项之和。）
//0、1、1、2、3、5、8、13、21、34、……
func Test_exercise_fibonacci_closure(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), "、")
	}
}

func fibonacci() func() int {
	//返回一个“返回int的函数”
	var num1, num2 int = -1, 1
	return func() int {
		num1, num2 = num2, num1+num2
		return num2
	}
}
