package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"

	"github.com/lukas-sgx/corel/pkg/utils"
)

func handleConnection(client net.Conn) {
	var data = bufio.NewReader(client);

	fmt.Println(utils.SetBlue("[INFO]") + " New agent handshake " + client.RemoteAddr().String())
	defer client.Close()

	for {
		line, err := data.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Print(line)
	}
	fmt.Println(utils.SetBlue("[INFO]") + " Agent " + client.RemoteAddr().String() + " lost connection")
}

func WaitConnection(tcp net.Listener) {
	for {
		client, err := tcp.Accept()
		
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(client)
	}
}