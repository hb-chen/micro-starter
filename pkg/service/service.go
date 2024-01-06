package service

import (
	"github.com/urfave/cli/v2"
	"micro.dev/v4/cmd"
	"micro.dev/v4/service"
	"micro.dev/v4/service/profile"
)

// New returns a new Micro Service
func New(opts ...service.Option) *service.Service {
	// setup micro, this triggers the Before
	// function which parses CLI flags.
	c := cmd.New(
		cmd.Service(),
		cmd.Flags(
			&cli.StringFlag{
				Name:    "address",
				Usage:   "Set the micro service address",
				EnvVars: []string{"MICRO_SERVICE_ADDRESS"},
				Value:   ":0",
			},
		),
		cmd.Before(func(ctx *cli.Context) error {
			addr := ctx.String("address")
			opts = append(opts, service.Address(addr))
			return nil
		}),
	)

	err := c.Run()
	if err != nil {
		return nil
	}

	// setup auth
	profile.SetupAccount(nil)

	// return a new service
	svc := &service.Service{}
	svc.Init(opts...)
	return svc
}
