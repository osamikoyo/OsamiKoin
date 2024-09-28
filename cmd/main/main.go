package main

import (
	"osamikoin/internal/server"
)

func main() {
	var cher chan error
	var ch chan string
	go server.TCPserver(cher, ch)
	s := server.New()
	s.Run()

	select {}
}
