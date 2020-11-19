package learngo

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func Test_element(t *testing.T) {
	fmt.Println("标识符(identifier): 以字母a~z(大小写均可)或下划线_开头，后面可以是多个字母、下划线和数字。不允许使用标点符号，例如@、$、%等一系列符号")
	fmt.Println("关键词(keyword):")
	fmt.Println("运算符(operator):")
	fmt.Println("分隔符(delimiter):")
	fmt.Println("字面量(literal):")
}

func Test_element_identifier(t *testing.T) {
	fmt.Println("数据类型：bool、byte、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、float32、float64、uintptr、string、rune、error、complex64、complex128")
	fmt.Println("内建函数名：append、cap、close、complex, copy、delete、imag、len、make、new、panic、print、println、real、recover")
	fmt.Println("其它标识符：iota、nil、_")

	fmt.Println("byte // uint8 的别名")
	fmt.Println("rune // int32 的别名")
	fmt.Println("uintptr:用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。")
}

func Test_element_keyword(t *testing.T) {
	fmt.Println("包导入与声明：import、package")
	fmt.Println("程序实体声明与定义：var、type、func、interface、map、struct、chan、const")
	fmt.Println("流程控制：if、continue、for、return、go、case、goto、switch、select、else、break、default、defer、fallthrough、range")
}

func Test_element_literal(t *testing.T) {
	fmt.Println("Struct(结构体)、Array(数组)、Slice(切面)、Map(字典)")
}

func Test_element_other(t *testing.T) {
	fmt.Println("_ ：空标识符")
	fmt.Println(":= ：一种简写，表示函数内部变量声明并赋值")
	fmt.Println(" ++ 和 -- 是语句，不是表达式，没有运算符优先级之说")
}

func Test_const(t *testing.T) {
	fmt.Println("常量：const {变量名} type = {赋值}")
	fmt.Println("显示定义：const h string = \"hello\"")
	fmt.Println("隐示定义：const w = \"world\"")
	fmt.Println("反斜杠\\，可以在定义常量的表达式中，作为跨行链接符，与shell命令一样")
	fmt.Println("整体数字后缀U（unsigned）、L(long)")
	fmt.Println("转义字符 \\n(新行)\\r(回车)")
	fmt.Println("常量还可以用作枚举，使用iota关键词实现枚举，第一个常量被默认设置为0，后续的常量默认设置为它上面那个常量的值，iota则递增，所以可是实现枚举")
	const Pi = 3.1415926
	const 人数, 费用, 班级 = 40, 200.01, "一班"
	const (
		Monday, Tuesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	const (
		Connected    = 0
		Disconnected = 1
		Unknown      = 2
	)

	const (
		a       = iota             //a==0
		b                          //b==1,隐示使用iota关键词，实际等同于b==iota
		c                          //c==2,隐示等同于c=iota
		d, e, f = iota, iota, iota //d==3,e==3,f==3同一行值相同，此处不能只写一个iota
		g       = iota             //g==4
		h       = "h"              //h=="h",单独赋值，iota依旧递增为5
		i                          //i=="h",默认使用上面的赋值，iota依旧递增为6
		j       = iota             //j==7
		k       = iota + 50        //k==7+1+50,iota依旧递增为8,并+50
	)
	const z = iota //每个单独定义的const常量中，iota都会重置，此时iota==0
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, z)
}

func Test_var(t *testing.T) {
	fmt.Println("变量：var {变量名} type = {赋值}")
	fmt.Println("静态声明：var {变量名} type ，动态声明：var {变量名} = {赋值}")
	fmt.Println("推导声明(只能在函数体内)：{变量名} := {赋值}")
	fmt.Println("两数交换：a,b=b,a")
	fmt.Println("变量的类型也可以在运行时实现自动推断")
	fmt.Println("局部变量定义了，就必须使用，即使赋值也不通过编译检查，但可以用空标识符(_=a)来忽略编译检查")
	fmt.Println("init函数：特殊的函数，类似类的构造方法，它会在包完成初始化后自动执行,执行优先级高于main方法，并不能手动调用init函数，每个源文件有且只有一个init函数，初始化过程会根据包的依赖关系按顺序单线程执行")
	var a, b, c = "string", 1, true
	var (
		d, e int
		f    bool
		g    string
	)
	var x int

	var (
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)

	num, _ := math.Modf(123.465)

	fmt.Println(a, b, c, d, e, f, g, x)
	fmt.Println(HOME, USER, GOROOT)
	fmt.Println(Pi, Pi*Pi)
	fmt.Println(num)
}

var Pi float64

func init() {
	Pi = 4 * math.Atan(1) //在init函数中计算 Pi 的值
}
