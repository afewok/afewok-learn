package env_test

import (
	"fmt"
	"testing"
)

func Test_env(t *testing.T) {
	fmt.Println("GOROOT：表示Go在你的电脑上的安装位置")
	fmt.Println("GOBIN：表示编译器和链接器的安装位置，默认是 GOROOT/bin")
	fmt.Println("GOHOSTARCH：表示本机处理器架构（不设置，则与GOARCH一致）")
	fmt.Println("GOHOSTOS：表示本地机器（不设置，则与GOOS一致）")
	fmt.Println("GOARCH：表示目标机器的处理器架构（386、amd64、arm）")
	fmt.Println("GOOS：表示目标机器的操作系统（darwin、freebsd、linux、windows）")

	fmt.Println("GOARM：专门针对基于ARM架构的处理器，可以是5 或 6。默认为6")
	fmt.Println("GOMAXPROCS：用于设置应用程序可使用的处理器个数与核数")
	fmt.Println("GOPATH（已废弃）：工作目录，可以有多个，类似于工作空间的概念，包含\\bin（可执行文件）、\\pkg（包文件）、\\src（源码文件）目录")

}
