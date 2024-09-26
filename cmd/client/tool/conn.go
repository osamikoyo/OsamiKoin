package tool

import (
	"net"
	"sync"
)

func ConnectionRouting(ch chan string, che chan error, listner net.Listener, wg *sync.WaitGroup) {
	for {
		conn, err := listner.Accept()
		if err != nil {
			che <- err
			wg.Done()
		}
		conn.Write([]byte("message"))
		conn.Close()
	}
}