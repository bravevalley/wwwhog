package cmd

import (
	"fmt"
	"os/exec"
)


func Getlisting(username string) (string, error) {
	
	cmd, err := exec.Command("/usr/local/bin/fetch_backup.sh").Output()

	if err != nil {
		return "", fmt.Errorf("%s error: %v\n", "/usr/local/bin/fetch_backup.sh", err)
	}
	
	return string(cmd), nil

	
}

