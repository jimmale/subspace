package listener

import "github.com/lucas-clemente/quic-go"

// Singleton Pattern
var currentListener *Listener

func GetListener() *Listener {
	if currentListener == nil {
		currentListener = newListener()
	}
	return currentListener
}

type Listener struct {
	qListener quic.Listener
}

func newListener() *Listener {
	return nil // TODO
}

func (l *Listener) Close() {

	l.qListener.Close()
	currentListener = nil
}
