package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"
)

func TestReadData(t *testing.T) {
	server, client := net.Pipe()
	defer func () {
		if server.Close() != nil {
			log.Fatal("Error on close pipe")
		}
	}()
	defer func () {
		if client.Close() != nil {
			log.Fatal("Error on close pipe")
		}
	}()

	go func() {
		for range 2 {
			_, err := fmt.Fprint(client, "KEEPALIVE\n")
			if err != nil {
				log.Fatal("Failed to assign value")
			}
		}
		_, err := fmt.Fprint(client, "Hello World\n")
		if err != nil {
			log.Fatal("Failed to assign value")
		}
	}()

	reader := bufio.NewReader(server)
	result, err := readData(reader, server)

	if err != nil {
		t.Fatalf("Erreur inattendue : %v", err)
	}
	if result != "Hello World" {
		t.Errorf("Attendu 'Hello World', obtenu '%s'", result)
	}
}