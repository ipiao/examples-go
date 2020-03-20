package net

import (
	"net"
	"testing"
)

func TestNetListenZeroAddr(t *testing.T) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	t.Log(ln.Addr())
	// Output:
	// [::]:random-port
}
