package server

import (
	"log"
	"net"
	"strings"

	"com.go54/rbso/cmd"
)

var (
	Get_req = "GET"
 	Keep_req = "KEEP"
 	hstnme string
 	out string
)

// Startserver starts the server and listens on the specified port
func Startserver(port string) {
	listener, err := net.Listen("tcp", port)
	
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()
	log.Printf("Server listening on %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Couldn't accept connection")
			continue
		}
		go handleconn(conn)
	}
}

// handleconn process all the incoming connection and route them to the right function
func handleconn(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 1024)

	_, err := conn.Read(b)
	if err != nil {
		log.Fatal("Failed to read data from connection")
	}

	query := strings.Fields(string(b))


	if len(query) > 4 {
		log.Fatalf("Wrong!") // TODO: Change later
	}




	switch request := strings.ToUpper(query[1]); request {
	case Get_req:
		var un = strings.ToLower(query[2])
		hstnme = strings.ToLower(query[3])

		out, err = cmd.Getlisting(un, hstnme)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s GET %v %v - success", conn.RemoteAddr(), un, hstnme)
	case Keep_req:
		var metadata = strings.ToLower(query[3])
		hstnme = strings.ToLower(query[2])

		out, err = cmd.Keepdata(hstnme, metadata)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s KEEP %v %v - success", conn.RemoteAddr(), hstnme, metadata)
	default:
		log.Fatalf("Action \"%v\" unknown; kindly use GET or KEEP", request)

	}

	conn.Write([]byte(out))
}