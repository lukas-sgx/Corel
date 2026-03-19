package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"encoding/json"

	"github.com/lukas-sgx/corel/pkg/utils"
)

func readData(data *bufio.Reader, client net.Conn) (string, error) {
	for {
		line, err := data.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf(utils.SetRed("[ERROR]") + " connection read error from %s: %v", client.RemoteAddr().String(), err)
			break
		}
		line = strings.TrimSpace(line)
		if line != "KEEPALIVE" {
			return line, nil
		}
	}
	return "", errors.New("invalid read or lost connection")
}

func fetchClientData(data *bufio.Reader, client net.Conn, dataAuth utils.Auth) {
	for {
		response, err := readData(data, client)
		if err != nil {
			fmt.Println(utils.SetBlue("[INFO]") + " Agent " + client.RemoteAddr().String() + " lost connection")
			break
		}

		err = json.Unmarshal([]byte(response), &dataAuth)
		if err != nil {
			log.Fatal(utils.SetRed("ERROR"), " fatal reading json")
		}

		fmt.Println(utils.SetBlue("[INFO] ") + dataAuth.Identity + " assign -> " + client.RemoteAddr().String())
	}
}

func handleConnection(client net.Conn) {
	var data = bufio.NewReader(client);
	var dataAuth utils.Auth

	fmt.Println(utils.SetBlue("[INFO]") + " New agent handshake " + client.RemoteAddr().String())
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf(utils.SetRed("[ERROR]") + " failed to close connection: %v", err)
		}
	}()

	fetchClientData(data, client, dataAuth)
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