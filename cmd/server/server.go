package server

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var (
	listenAddr string
	port int
	meshSync bool
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start Corel server C2C",
	Long:    `Start Corel server C2C listener, handle incoming connections.`,
	GroupID: "modules",
	Run: serverRun,
}

func serverRun(cmd *cobra.Command, args []string) {
	tcpListen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", listenAddr, port))
	
	if err != nil {
		log.Fatal(err)
	}
	defer tcpListen.Close()

	initHeader()
	isMeshSync(meshSync)
	listenerInfo(listenAddr, port)
	
	waitConnection(tcpListen)
}

func init() {
	ServerCmd.Flags().StringVarP(&listenAddr, "LHOST", "l", "0.0.0.0", "Address to listen on")
	ServerCmd.Flags().IntVarP(&port, "LPORT", "p", 8080, "Port to listen on")
	ServerCmd.Flags().BoolVarP(&meshSync, "MSYNC", "m", true, "Use mesh synchronization")
}
