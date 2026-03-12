package agent

import (
	"fmt"
	"log"
	"net"
)

type Agent struct {
	RemoteAddr string
	RemotePort int
	MeshSync bool
	Version string
	Identity string
}

func launchHandshake(tcp net.Conn) {
	defer tcp.Close()
}

func Peer(agent Agent) {
	tcpConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", agent.RemoteAddr, agent.RemotePort))

	if err != nil {
		log.Fatal(err)
	}
	launchHandshake(tcpConn)
}