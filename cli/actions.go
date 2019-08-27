package cli

import (
	"fmt"
	"os"

	"github.com/kaz/keymold/otp"
	"github.com/lox/go-touchid"
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

	secret, err := otp.LoadSecret(keyName)
	if err != nil {
		return fmt.Errorf("loading secret failed: %v", err)
	}

	if v, ok := secret.Options[FLAG_DISABLE_TOUCH_ID]; !ok || !v.(bool) {
		if ok, err = touchid.Authenticate("Generate OTP Code"); err != nil {
			return fmt.Errorf("TouchID authentication failed with error: %v", err)
		} else if !ok {
			return fmt.Errorf("TouchID authentication failed")
		}
	}

	code, err := secret.GetCode()
	if err != nil {
		return fmt.Errorf("generating code failed: %v", err)
	}

	fmt.Print(code)
	return nil
}
