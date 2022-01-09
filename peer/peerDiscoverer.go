package peer

import "net"

type PeerDiscoverer interface {
	GetPeers() ([]PeerConnectionInfo, error)
}

type PeerConnectionInfo struct {
	Address net.IP
	Port    uint
}
