package ssh

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/kevinburke/ssh_config"
)

func parseDest(dest string) (string, string, string, error) {
	subm := regexp.MustCompile("^(?:(.+?)@)?(.+?)(?::(\\d+?))?$").FindSubmatch([]byte(dest))
	if subm == nil {
		return "", "", "", fmt.Errorf("invalid destination string")
	}

	return string(subm[1]), string(subm[2]), string(subm[3]), nil
}
func parseDestWithFile(dest string) (string, string, error) {
	userName, hostName, port, err := parseDest(dest)
	if err != nil {
		return "", "", err
	}

	if userName == "" {
		userName = ssh_config.Get(hostName, "User")
	}
	if userName == "" {
		userName = ssh_config.Default("User")
	}
	if userName == "" {
		currentUser, err := user.Current()
		if err != nil {
			return "", "", fmt.Errorf("checking user failed: %v", err)
		}
		userName = currentUser.Name
	}

	if port == "" {
		port = ssh_config.Get(hostName, "Port")
	}
	if port == "" {
		port = ssh_config.Default("Port")
	}
	if port == "" {
		port = "22"
	}

	configuredHostName := ssh_config.Get(hostName, "HostName")
	if configuredHostName != "" {
		hostName = configuredHostName
	}

	hostName += ":" + port
	return userName, hostName, nil
}
