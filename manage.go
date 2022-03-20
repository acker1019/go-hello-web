package main

import (
	"os"

	"ack/go-hello-web/core"
	"ack/go-hello-web/framework"
)

func main() {
	args := os.Args
	cmds := args[1:]
	var core core.Core
	framework.RouteCommand(&core, cmds)
}
