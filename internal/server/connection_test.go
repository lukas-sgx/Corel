package server

import (
	"bufio"
	"testing"
	"net"
	"fmt"
)

func TestReadData(t *testing.T) {
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	go func() {
		fmt.Fprint(client, "KEEPALIVE\n")
		fmt.Fprint(client, "KEEPALIVE\n")
		fmt.Fprint(client, "Hello World\n")
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