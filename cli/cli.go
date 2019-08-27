package cli

import (
	"os"

	"github.com/urfave/cli"
)

const (
	FLAG_DISABLE_TOUCH_ID = "disable-touch-id"
)

func Start() error {
	app := cli.NewApp()

	app.Name = "keymold"
	app.Version = "0.0.1" // TODO: Specify on build process
	app.Description = "OTP generator, works on command line."

	app.Commands = []cli.Command{
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "add new OTP secret",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  FLAG_DISABLE_TOUCH_ID,
					Usage: "allow generating OTP without TouchID authentication (insecure)",
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
	}

	return app.Run(os.Args)
}
