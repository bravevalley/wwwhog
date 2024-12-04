package cmd

import (
	"fmt"
	"os/exec"
)

func Getlisting(username, server_id string) (string, error) {

	cmd, err := exec.Command("/usr/local/bin/fetch_backup.sh", "username", "server_id").Output()

	if err != nil {
		return "", fmt.Errorf("%s error: %v", "/usr/local/bin/fetch_backup.sh", err)
	}

	return string(cmd), nil

}
