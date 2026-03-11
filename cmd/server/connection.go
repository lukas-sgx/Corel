package server

import (
	"log"
	"net"
	"fmt"
)

func handleConnection(client net.Conn) {
	fmt.Println(client.LocalAddr().String())
	defer client.Close()
}

func waitConnection(tcp net.Listener) {
	for {
		client, err := tcp.Accept()
		
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(client)
	}
}