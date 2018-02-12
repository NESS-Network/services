package cli

import "github.com/urfave/cli"

var (
	rpcaddr  = new(string)
	endpoint string
)

// App is a cli app
var App = cli.App{
	Commands: []cli.Command{
		{
			Name: "btc",
			Subcommands: []cli.Command{
				generateAddrCMD(),
				generateKeyPairCMD(),
				checkBalanceCMD(),
			},
			Before: func(c *cli.Context) error {
				endpoint = "c2cx"
				return nil
			},
		},
	},
	Flags: []cli.Flag{
		cli.StringFlag{Name: "rpc", Destination: rpcaddr, Value: "localhost:12345"},
	},
}
