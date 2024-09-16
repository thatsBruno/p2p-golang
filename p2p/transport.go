package p2p

// Peer is an interface that represents the remote node.
type Peer interface{}

// Transport is anything that handles the communication
// between the nodes in the network. Thi can be of
// form (TCP, UDP, websockets, ...)
type Transport interface {
	// 25:00
	ListenAndAccept() error
}
