package peer

import "net"
import "github.com/lucas-clemente/quic-go"

type Peer struct {
	Session quic.Session
}

func Connect(ip net.IP, port uint) (*Peer, error) {

}
