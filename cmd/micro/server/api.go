package server

import (
	"github.com/hb-chen/micro-starter/cmd/micro/service/api"
	"github.com/urfave/cli/v2"
)

func runAPI(ctx *cli.Context, wait chan bool) error {
	apiCmd := api.NewAPI(wait)
	return apiCmd.Run(ctx)
}
