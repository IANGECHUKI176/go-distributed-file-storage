package p2p

import (
	"fmt"
	"net"
	"sync"
)

//remote node over a tcp est conn
type TCPPeer struct{
	//underlying conn of the peer
	conn net.Conn
	//if we dial a connection outbound->true
	//if we accept a connection outbound ->false
	outbound bool
}

func NewTCPpeer(conn net.Conn, outbound bool) *TCPPeer{
	return &TCPPeer{
		conn:conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct{
	listenAddress string

	HandeshakeFunc HandeshakeFunc
	Decoder Decoder
}
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	mu    sync.RWMutex
	peers map[net.Addr]Peer
}
//public func at top and private func at bottom

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		
	}
}

func (t *TCPTransport) ListenAndAccept() error{
	var err error
	t.listener,err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop(){
	for {
		conn ,err := t.listener.Accept()
		if err != nil {
			fmt.Printf("tcp  accepr error: %s\n",err)
		}
		go t.handleConn(conn)
	}
}
type Temp struct {}
func (t *TCPTransport) handleConn(conn net.Conn){
	peer := NewTCPpeer(conn,true)

	
	if err := t.HandeshakeFunc(peer); err != nil {
		conn.Close()
		return
	}
	//read loop
	msg := &Temp{}
	for {
		if err := t.Decoder.Decode(conn,msg); err != nil{
			fmt.Printf("tcp read error: %s\n",err)
			continue
		}

	}

}


