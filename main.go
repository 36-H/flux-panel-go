package main

import (
	_ "flux-panel-go/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"flux-panel-go/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
