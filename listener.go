package nalu

import (
	"net"
)

// Listens for new connections
type listener chan net.Conn

// Functions to satisfy interface net.Listener

func (l listener) Accept() (c net.Conn, err error) { return <-l, nil }
func (l listener) Close() (err error)              { return nil }
func (l listener) Addr() (a net.Addr)              { return nil }
