package main

import (
	_ "Shop/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"Shop/internal/cmd"

	_ "Shop/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
