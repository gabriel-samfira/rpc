package rpc

import (
	"net"

	"gopkg.in/natefinch/npipe.v2"
)

func dialAddress(network, address string) (net.Conn, error) {
	switch network {
	case "npipe":
		return npipe.Dial(address)
	default:
		return net.Dial(network, address)
	}
}
