package main

import (
	"fmt"

	"goframe-study/cmd"

	"github.com/gogf/gf/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
	fmt.Println(1)
}
