package cli

import (
	"fmt"
	"os"

	"github.com/kaz/keymold/otp"
	"github.com/kaz/keymold/ssh"
	"github.com/urfave/cli"
)

func AddSecret(context *cli.Context) error {
	keyName := context.Args().Get(0)
	if keyName == "" {
		keyName = DEFAULT_KEY_NAME
	}

	secretValue := ""
	fmt.Fprint(os.Stderr, "Input your secret key: ")
	if _, err := fmt.Scanln(&secretValue); err != nil {
		return fmt.Errorf("reading secret failed: %v", err)
	}

	secret := otp.NewSecret(keyName, secretValue)
	secret.Options[FLAG_DISABLE_TOUCH_ID] = context.Bool(FLAG_DISABLE_TOUCH_ID)

	if err := secret.Save(); err != nil {
		return fmt.Errorf("saving secret failed: %v", err)
	}

	return nil
}

func GetCode(context *cli.Context) error {
	keyName := context.Args().Get(0)
	if keyName == "" {
		keyName = DEFAULT_KEY_NAME
	}

	code, err := getCode(keyName, "Get OTP")
	if err != nil {
		return fmt.Errorf("OTP generation failed: %v", err)
	}

	fmt.Print(code)
	return nil
}

func CreateProxy(context *cli.Context) error {
	keyName := context.Args().Get(0)
	if keyName == "" {
		keyName = DEFAULT_KEY_NAME
	}

	code, err := getCode(keyName, "Connect SSH server")
	if err != nil {
		return fmt.Errorf("OTP generation failed: %v", err)
	}

	return ssh.Pipe(context.String(FLAG_BASTION), context.String(FLAG_TARGET), code)
}
