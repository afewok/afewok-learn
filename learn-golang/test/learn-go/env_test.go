package testing

import (
	"fmt"
	"testing"
)

func Test_env_config(t *testing.T) {
	fmt.Println("GOROOT：表示Go在你的电脑上的安装位置")
	fmt.Println("GOBIN：表示编译器和链接器的安装位置，默认是 GOROOT/bin")
	fmt.Println("GOHOSTARCH：表示本机处理器架构（不设置，则与GOARCH一致）")
	fmt.Println("GOHOSTOS：表示本地机器（不设置，则与GOOS一致）")
	fmt.Println("GOARCH：表示目标机器的处理器架构（386、amd64、arm）")
	fmt.Println("GOOS：表示目标机器的操作系统（darwin、freebsd、linux、windows）")

	fmt.Println("GOARM：专门针对基于ARM架构的处理器，可以是5 或 6。默认为6")
	fmt.Println("GOMAXPROCS：用于设置应用程序可使用的处理器个数与核数")
	fmt.Println("GOPATH（已废弃）：工作目录，可以有多个，类似于工作空间的概念，包含\\bin（可执行文件）、\\pkg（包文件）、\\src（源码文件）目录")

	fmt.Println("#Mac上配置示例 go ")
	fmt.Println("export GOROOT=/usr/local/go")
	fmt.Println("export GOPATH=/Users/{user}/.go")
	fmt.Println("export GOBIN=$GOPATH/bin")
	fmt.Println("export PATH=$GOPATH/bin:$GOROOT/bin:$PATH")
}

func Test_command_line(t *testing.T) {
	fmt.Println("go mod：模块化（管理软件的版本号）")
	fmt.Println("go clean：删除对象文件")
	fmt.Println("go fmt：格式化源码包的代码")
	fmt.Println("go tool：运行一个指定的go tool")
	fmt.Println("go vet：检查源码包中可能出现的错误")

	fmt.Println("go get 下载并安装指定的包与依赖")
	fmt.Println("go install：编译并安装指定的包与依赖")
	fmt.Println("go build：编译源代码包和依赖")

	fmt.Println("go run：编译并运行Go程序")
	fmt.Println("go test：测试一个源码包")

	fmt.Println("go list：打印指定源码包的信息")
	fmt.Println("go version：打印输出Go环境版本")
	fmt.Println("go doc：显示Go包或程序实体的文档")
	fmt.Println("go env：打印输出Go语言环境的变量信息")
	fmt.Println("go generate：通过扫描Go源码中的go：generate注释来识别要运行的常规命令")
	fmt.Println("go bug：bug提交程序")
	fmt.Println("go fix：修复程序")
}
