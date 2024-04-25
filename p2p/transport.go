package p2p

//interface that represents the remote node
type Peer interface{}

//anything that handles communication 
//between peers in the network (TCP,UDP,....)
type Transport interface{
	ListenAndAccept() error
}