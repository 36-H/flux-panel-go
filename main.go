package main

import (
	_ "flux-panel-go/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"flux-panel-go/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
