package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() (Peer, error) {
	for {

		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting TCP connection: %v\n", err)
			return nil, err
		}

		go t.handleConn(conn)
	}

}

func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf("Accepted TCP connection from %v\n", conn.RemoteAddr())
}
