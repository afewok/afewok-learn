package main

import (
	_ "learn-golang/boot"
	_ "learn-golang/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
