package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"com.go54/rbso/cmd"
)


var (
	Get_req = "GET"
 	Keep_req = "KEEP"
 	hstnme string
 	out string
 	err error
)

func main() {

	if len(os.Args) > 4 {
		log.Fatalf("Wrong!") // TODO: Change later
	}




	switch request := strings.ToUpper(os.Args[1]); request {
	case Get_req:
		var un = strings.ToLower(os.Args[2])
		hstnme = strings.ToLower(os.Args[3])

		out, err = cmd.Getlisting(un, hstnme)
		if err != nil {
			log.Fatal(err)
		}
	case Keep_req:
		var metadata = strings.ToLower(os.Args[3])
		hstnme = strings.ToLower(os.Args[2])

		out, err = cmd.Keepdata(hstnme, metadata)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Action \"%v\" unknown; kindly use GET or KEEP", request)

	}


	fmt.Println(out)
}
