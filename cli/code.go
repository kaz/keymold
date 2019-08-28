package cli

import (
	"fmt"

	"github.com/kaz/keymold/otp"
	"github.com/lox/go-touchid"
)

func getCode(keyName string, reason string) (string, error) {
	secret, err := otp.LoadSecret(keyName)
	if err != nil {
		return "", fmt.Errorf("loading secret failed: %v", err)
	}

	if v, ok := secret.Options[FLAG_DISABLE_TOUCH_ID]; !ok || !v.(bool) {
		if ok, err = touchid.Authenticate(reason); err != nil {
			return "", fmt.Errorf("TouchID authentication failed with error: %v", err)
		} else if !ok {
			return "", fmt.Errorf("TouchID authentication failed")
		}
	}

	code, err := secret.GetCode()
	if err != nil {
		return "", fmt.Errorf("generating code failed: %v", err)
	}

	return code, nil
}
