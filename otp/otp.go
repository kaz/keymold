package otp

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func (s *secret) GetCode() (string, error) {
	return totp.GenerateCode(s.Secret, time.Now())
}
