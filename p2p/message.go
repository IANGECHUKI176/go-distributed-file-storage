package p2p

import "net"

//message hold an arbitrary data that can be sent
//over between two nodes in a network
type RPC struct{
	From net.Addr
	Payload []byte
	
}
