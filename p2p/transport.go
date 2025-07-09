package p2p

// Peer is an interfact that represetns the remote node
type Peer interface {

}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the 
// form of TCP, UDP, websockets, WebRTC, etc.
type Transport interface {
	ListenAndAccept() error
}

