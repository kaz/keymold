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
		return fmt.Errorf("key_name is empty (see --help)")
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
		return fmt.Errorf("key_name is empty (see --help)")
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
		return fmt.Errorf("key_name is empty (see --help)")
	}

	bastionDest := context.Args().Get(1)
	if bastionDest == "" {
		return fmt.Errorf("bastion_dest is empty (see --help)")
	}

	targetDest := context.Args().Get(2)
	if targetDest == "" {
		return fmt.Errorf("target_dest is empty (see --help)")
	}

	code, err := getCode(keyName, "Connect SSH server")
	if err != nil {
		return fmt.Errorf("OTP generation failed: %v", err)
	}

	return ssh.Pipe(bastionDest, targetDest, code)
}
