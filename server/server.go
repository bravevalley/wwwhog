package server

import (
	"fmt"
	"net"
	"strings"

	"com.go54/rbso/cmd"
	"com.go54/rbso/log"
)

var (
	Get_req = "GET"
 	Keep_req = "KEEP"
 	hstnme string
 	out string
)

// Startserver starts the server and listens on the specified port
func Startserver(port string) {
	listener, err := net.Listen("tcp4", port)
	
	if err != nil {
		logah.Logger.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()
	logah.Logger.Printf("Server listening on %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logah.Logger.Println("Couldn't accept connection")
			continue
		}
		go handleconn(conn)
	}
}

// handleconn process all the incoming connection and route them to the right function
func handleconn(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)

	_, err := conn.Read(b)
	if err != nil {
		logah.Logger.Println("Failed to read data from connection")
		return
	}

	cleanedStr := strings.Trim(string(b), "\x00")
	query := strings.Split(cleanedStr, " ")


	if len(query) > 4 {
		logah.Logger.Printf("%v: wrong query %v", conn.RemoteAddr().String(), strings.Join(query, " ")) // TODO: Change later
		return
	}




	switch request := strings.ToUpper(query[0]); request {
	case Get_req:
		var un = strings.ToLower(query[1])
		hstnme = strings.ToLower(query[2])

		out, err = cmd.Getlisting(un, hstnme)
		if err != nil {
			logah.Logger.Println(err)
			out = fmt.Sprintln("Error: Get backup information")
			break
		}

		if out == "-1" {
			out = fmt.Sprintln("No backup :(")
			logah.Logger.Printf("%s GET %v %v - no backup", conn.RemoteAddr().String(), un, hstnme)
			break
		}

		logah.Logger.Printf("%s GET %v %v - success", conn.RemoteAddr().String(), un, hstnme)

	case Keep_req:
		var metadata = strings.ToLower(query[2])
		hstnme = strings.ToLower(query[1])

		out, err = cmd.Keepdata(hstnme, metadata)
		if err != nil {
			logah.Logger.Println(err)
			out = fmt.Sprintf("Error: Couldn't move %v to host\n", metadata)
			break
		}

		logah.Logger.Printf("%s KEEP %v %v - success", conn.RemoteAddr().String(), hstnme, metadata)
	default:
		logah.Logger.Println(err)

	}

	dt := strings.Trim(out, "\n")

	conn.Write([]byte(dt))
}