package quic

import (
	"context"
	"fmt"
	"net"
	"time"
)

type TracedConn struct {
	net.PacketConn
	Ctx context.Context
}

func (t *TracedConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	start := time.Now()
	from, a, err := t.PacketConn.ReadFrom(p)
	fmt.Println("read", a, len(p), time.Since(start))
	return from, a, err
}

func (t *TracedConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	start := time.Now()
	n, err = t.PacketConn.WriteTo(p, addr)
	fmt.Println("write", addr, len(p), time.Since(start))
	return n, err
}
