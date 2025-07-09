package p2p

import "errors"

// ErrrInvalidHandshake is returned if the handshake between 
// the local and remote node could not be established

var ErrInvalidHandshake = errors.New("invalid handshake")
// HandshakeFunc is 
type HandshakeFunc func(Peer) error 

func NOPHandshakeFunc(Peer) error { return nil }
