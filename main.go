package main

import (
	_ "hello-gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"hello-gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
