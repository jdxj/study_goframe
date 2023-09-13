package main

import (
	_ "study_goframe/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"study_goframe/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
