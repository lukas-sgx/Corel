package agent

import (
	"fmt"
	"log"
	"net"

	"github.com/lukas-sgx/corel/pkg/utils"
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
	for {
		tcp.Write([]byte("KEEPALIVE"))
	}
}

func Peer(agent Agent) {
	tcpConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", agent.RemoteAddr, agent.RemotePort))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println(utils.SetBlue("[INFO]") + " Launch Hanshake on server...")
	launchHandshake(tcpConn)
}