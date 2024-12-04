package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// Getlisting takes the username and the server hostname and calls the server native script to retrieve
// the user's available backup and a nil error if no error occured
func Getlisting(username, hostname string) (string, error) {

	log.Printf("user: %q host: %q\n", username, hostname)

	cmd, err := exec.Command("/usr/local/bin/fetch_backup.sh", username, hostname).Output()

    if err != nil {
        log.Fatal(err)
    }

	prettyOut, err := PrettyString(string(cmd))
	if err != nil {
		return "", err
	}

	return prettyOut, nil
}



// Keepdata takes the hostname destination server and filepath of the backup
// then moves the backup from the source filepath to the destination
func Keepdata(dst, metadata string) (string, error) {

	cmd, err := exec.Command("/usr/local/bin/keep_backup.sh", dst, metadata).Output()

	if err != nil {
		return "", fmt.Errorf("%s failed to move backup to %s: %v", "/usr/local/bin/keep_backup.sh", dst, err)
	}

	return string(cmd), nil
}





// PrettyString takes json encoded string and return an indented string version of the json and a nil if 
// parsng was successful
func PrettyString(str string) (string, error) {

	var prettyJSON bytes.Buffer

	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}
