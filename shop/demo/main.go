package main

import (
	_ "demo/internal/packed"

	_ "demo/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"demo/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}
