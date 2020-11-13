package testing

import (
	"fmt"
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
	fmt.Println("数据类型：bool、byte、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、float32、float64、uintptr、string")
	fmt.Println("内建函数名：append、cap、close、complex、copy、delete、imag、len、make、new、panic、print、println、real、recover")
	fmt.Println("其它标识符：iota、nil、_")
}

func Test_element_keyword(t *testing.T) {
	fmt.Println("包导入与声明：import、package")
	fmt.Println("程序实体声明与定义：var、type、func、interface、map、struct、chan、const")
	fmt.Println("流程控制：if、continue、for、return、go、case、goto、switch、select、else、break、default、defer、fallthrough、range")
}

func Test_element_other(t *testing.T) {
	fmt.Println("_ ：空标识符")
	fmt.Println(":= ：")
}

func Test_const(t *testing.T) {
	fmt.Println("常量：const {变量名} type = {赋值}")

}

func Test_var(t *testing.T) {
	fmt.Println("变量：var {变量名} type = {赋值}")
}
