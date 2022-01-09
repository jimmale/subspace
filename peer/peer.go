package peer

import (
	"github.com/lucas-clemente/quic-go"
	log "github.com/sirupsen/logrus"
	"net"
)

type Peer struct {
	Session quic.Session
}

func (p *Peer) Disconnect() error {
	log.Tracef("Disconnecting from %s", p.Session.RemoteAddr().String())
	return p.Session.CloseWithError(0, "none")
}

func Connect(ip net.IP, port uint) (*Peer, error) {
	log.Tracef("Connecting to %s:%d", ip.String(), port)
	return nil, nil
}
