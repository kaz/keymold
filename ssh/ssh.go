package ssh

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func connect(user, addr, otp string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				if len(questions) != 1 {
					return nil, nil
				}
				return []string{otp}, nil
			}),
			ssh.PublicKeysCallback(func() ([]ssh.Signer, error) {
				sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
				if err != nil {
					return nil, fmt.Errorf("connecting to ssh-agent failed: %v", err)
				}

				signers, err := agent.NewClient(sock).Signers()
				if err != nil {
					return nil, fmt.Errorf("fetching key failed: %v", err)
				}

				return signers, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("ssh dial failed: %v", err)
	}

	return client, nil
}
