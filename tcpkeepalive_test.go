package tcpkeepalive

import (
	"net"
	"testing"
	"time"
)

func TestSetKeepAliveIdle(t *testing.T) {
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer ln.Close()
	go func() {
		c, err := ln.Accept()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer c.Close()
	}()

	c, err := net.Dial(ln.Addr().Network(), ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	if err := SetKeepAliveIdle(c.(*net.TCPConn), 5*time.Second); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := SetKeepAliveCount(c.(*net.TCPConn), 10); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := SetKeepAliveInterval(c.(*net.TCPConn), 2*time.Second); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
