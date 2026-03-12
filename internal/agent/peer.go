package agent

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/lukas-sgx/corel/pkg/utils"
)

type Agent struct {
	RemoteAddr string
	RemotePort int
	MeshSync   bool
	Version    string
	Identity   string
}

func keepAlive(tcp net.Conn) {
	ticker := time.NewTicker(600 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		if _, err := tcp.Write([]byte("KEEPALIVE\n")); err != nil {
			break
		}
	}
}

func initDataSend(tcp net.Conn, data Agent) {
	dataAuth := utils.Auth{
		Identity: data.Identity,
	}

	payload, err := json.Marshal(dataAuth)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tcp.Write(append(payload, '\n')); err != nil {
		log.Fatal(err)
	}
}

func launchHandshake(tcp net.Conn, data Agent) {
	defer tcp.Close()
	initDataSend(tcp, data)
	keepAlive(tcp)
}

func Peer(agent Agent) {
	tcpConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", agent.RemoteAddr, agent.RemotePort))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println(utils.SetBlue("[INFO]") + " Launch Hanshake on server...")
	launchHandshake(tcpConn, agent)
}
