package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)



func TestServer(t *testing.T) {
	ddl := time.Now().Add(time.Second * 10)
	ctx, cancelCtx := context.WithDeadline(context.TODO(), ddl)
	defer cancelCtx()

	var dyla net.Dialer

	conn, err := dyla.DialContext(ctx, "tcp", "131.153.147.50:40411")
	if err != nil {
		t.Fatal(err)
	}
	
	defer conn.Close()

	_, err = conn.Write([]byte("get mubeenn1 shark"))
	if err != nil {
		t.Logf("Write operation failed: %v", err)
	}

	b := make([]byte, 1536)
	conn.Read(b)

	dt := strings.Trim(string(b), "\x00")
	fmt.Printf("%q, %d", dt, len(dt))

}