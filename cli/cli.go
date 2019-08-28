package cli

import (
	"os"

	"github.com/urfave/cli"
)

const (
	DEFAULT_KEY_NAME = "keymold_default_key_name"

	FLAG_DISABLE_TOUCH_ID = "disable-touch-id"

	FLAG_BASTION       = "bastion"
	FLAG_BASTION_SHORT = "b"
	FLAG_TARGET        = "target"
	FLAG_TARGET_SHORT  = "t"
)

var (
	Version = "dev-build"
)

func Start() error {
	app := cli.NewApp()

	app.Name = "keymold"
	app.Version = Version
	app.Usage = "OTP generator, works on command line."

	app.Commands = []cli.Command{
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "add new OTP secret",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  FLAG_DISABLE_TOUCH_ID,
					Usage: "Allow generating OTP without TouchID authentication. (insecure!)",
				},
			},
			ArgsUsage: "key_name",
			Action:    AddSecret,
		},
		{
			Name:      "get",
			ShortName: "g",
			Usage:     "generate OTP",
			ArgsUsage: "key_name",
			Action:    GetCode,
		},
		{
			Name:      "proxy",
			ShortName: "p",
			Usage:     "create SSH proxy tunnel",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     FLAG_BASTION_SHORT + "," + FLAG_BASTION,
					Usage:    "Destination of bastion server. example: [USER_NAME@]HOST_NAME[:PORT]",
					Required: true,
				},
				cli.StringFlag{
					Name:     FLAG_TARGET_SHORT + "," + FLAG_TARGET,
					Usage:    "Destination of target server. example: [USER_NAME@]HOST_NAME[:PORT]",
					Required: true,
				},
			},
			ArgsUsage: "key_name",
			Action:    CreateProxy,
		},
	}

	return app.Run(os.Args)
}
