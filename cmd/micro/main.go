package main

//go:generate ./scripts/generate.sh

import (
	"github.com/micro/micro/v3/cmd"

	// load packages so they can register commands
	_ "github.com/micro/micro/v3/cmd/service"

	_ "github.com/hb-chen/micro-starter/cmd/micro/server"
	_ "github.com/hb-chen/micro-starter/profile"
)

func main() {
	cmd.Run()
}
