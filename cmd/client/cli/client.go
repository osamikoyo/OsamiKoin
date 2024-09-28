package cli

import (
	"fmt"
	"log"
	"net"
)

func Client() {
	conn, err := net.Dial("tcp", "/conn")
	if err != nil {
		log.Println(err)
	}
	for {
		var command string
		_, err := fmt.Scanln(&command)
		if err != nil {
			log.Println(err)
		}
		conn.Write([]byte(command))
		buff := make([]byte, 1024)
		input, errs := conn.Read(buff)
		if errs != nil {
			log.Println(errs)
		}
		log.Println(string(input))
	}
}
