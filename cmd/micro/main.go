package main

//go:generate ./scripts/generate.sh

import (
	"micro.dev/v4/cmd"

	_ "github.com/hb-chen/micro-starter/cmd/micro/server"
	_ "github.com/hb-chen/micro-starter/cmd/micro/service"
	_ "github.com/hb-chen/micro-starter/cmd/micro/web"
	_ "github.com/hb-chen/micro-starter/profile"
)

func main() {
	cmd.Run()
}
