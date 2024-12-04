package main

import (
	"fmt"
	"log"

	"com.go54/rbso/cmd"
)


func main() {
	
	out, err := cmd.Getlisting("dovemini")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}