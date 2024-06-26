package main

import (
	"log"

	"github.com/IANGECHUKI176/go-distributed-file-storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HandeshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
		
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select{}
}
