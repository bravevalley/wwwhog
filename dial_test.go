package main

import (
	"context"
	"fmt"
	"net"
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

	_, err = conn.Write([]byte("get junewedd wghp3"))
	if err != nil {
		t.Logf("Write operation failed: %v", err)
	}

	b := make([]byte, 1 << 14)
	conn.Read(b)

	fmt.Println(string(b))

}