package otp

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/keybase/go-keychain"
)

type (
	secret struct {
		Name    string
		Secret  string
		Options map[string]interface{}
	}
)

const (
	KEYCHAIN_SERVICE = "keymold"
	KEYCHAIN_LABEL   = "keymold secret"
	KEYCHAIN_AG      = "com.narusejun.keymold"
)

func NewSecret(name, value string) *secret {
	return &secret{name, value, make(map[string]interface{})}
}

func LoadSecret(name string) (*secret, error) {
	data, err := keychain.GetGenericPassword(KEYCHAIN_SERVICE, name, KEYCHAIN_LABEL, KEYCHAIN_AG)
	if err != nil {
		return nil, fmt.Errorf("reading from keychain failed: %v", err)
	}

	s := &secret{}
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(s); err != nil {
		return nil, fmt.Errorf("deseliarization failed: %v", err)
	}

	return s, nil
}

func (s *secret) Save() error {
	buf := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buf).Encode(s); err != nil {
		return fmt.Errorf("seliarization failed: %v", err)
	}

	item := keychain.NewGenericPassword(KEYCHAIN_SERVICE, s.Name, KEYCHAIN_LABEL, buf.Bytes(), KEYCHAIN_AG)

	if err := keychain.AddItem(item); err != nil {
		if err != keychain.ErrorDuplicateItem {
			return fmt.Errorf("writing to keychain failed: %v", err)
		}

		if err := keychain.DeleteGenericPasswordItem(KEYCHAIN_SERVICE, s.Name); err != nil {
			return fmt.Errorf("existing item deletion failed: %v", err)
		}

		return s.Save()
	}

	return nil
}
