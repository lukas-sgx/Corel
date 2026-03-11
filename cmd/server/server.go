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
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start Corel server C2C",
	Long:    `Start Corel server C2C listener, handle incoming connections.`,
	GroupID: "modules",
	Run: serverRun,
}

func serverRun(cmd *cobra.Command, args []string) {
	_, err := net.Listen("tcp", fmt.Sprintf("%s:%d", listenAddr, port))

	InitHeader()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(SetGreen("[SCAN]") + " Listening for incoming beacons...")
	fmt.Println(SetGreen("[SCAN]") + " Global mesh synchronization: ACTIVE signification\n")
	fmt.Println(SetRed("[SERVER]") + " Server active on:")
	if (listenAddr == "127.0.0.1" || listenAddr == "0.0.0.0") {
		fmt.Printf(SetRed("[SERVER]") + "   - Internal: tcp://127.0.0.1:%d\n", port)
	}
	if (listenAddr != "127.0.0.1") {
		fmt.Printf(SetRed("[SERVER]") + "   - External: tcp://%s:%d\n", listenAddr, port)
	}
	fmt.Println()
	fmt.Println(SetGreen("[READY]") + "  Waiting for Client Handshakes...")
}

func init() {
	ServerCmd.Flags().StringVarP(&listenAddr, "LHOST", "l", "0.0.0.0", "Address to listen on")
	ServerCmd.Flags().IntVarP(&port, "LPORT", "p", 8080, "Port to listen on")
}
