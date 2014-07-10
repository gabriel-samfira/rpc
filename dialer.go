// +build !windows

package rpc

import (
	"net"
)

func dialAddress(network, address string) (net.Conn, error) {
	return net.Dial(network, address)
}
