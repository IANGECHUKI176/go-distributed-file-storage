package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTranport(t *testing.T){
	tcpOpts := TCPTransportOpts{
		ListenAddr: ":3000",
		HandeshakeFunc: NOPHandshakeFunc,
		Decoder: DefaultDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)

    assert.Equal(t, tr.ListenAddr, tcpOpts.ListenAddr)
	//server
	//tr.Listen
	assert.Nil(t,tr.ListenAndAccept())

	select{}

}