package quic

import (
	"context"
	"net"
)

type TracedConn struct {
	net.PacketConn
	Ctx context.Context
}

func (t *TracedConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	c := t.Ctx
	defer mon.Task()(&c)(&err)
	return t.PacketConn.ReadFrom(p)
}

func (t *TracedConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	c := t.Ctx
	defer mon.Task()(&c)(&err)
	return t.PacketConn.WriteTo(p, addr)
}
