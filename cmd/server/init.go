package server

import (
	"fmt"
	"github.com/lukas-sgx/corel/pkg/utils"
)

var header = 
utils.ClearScreen + utils.Blue +
"==============================\n" + utils.Bold +
"[ COREL  :: AUTONOMOUS AGENT ]\n" + utils.Reset + utils.Blue +
"[ MODE   ::           SERVER ]\n" +
"[ STATUS ::        LISTENING ]\n" +
"==============================\n" + utils.Reset

func initHeader() {
	fmt.Println(header)
	fmt.Println(utils.SetBlue("[SCAN]") + " Listening for incoming beacons...")
}

func isMeshSync(mesh bool) {
	if (mesh == true) {
		fmt.Println(utils.SetBlue("[SCAN]") + " Global mesh synchronization: ACTIVE signification\n");
	} else {
		fmt.Println()
	}
}

func listenerInfo(listenAddr string, port int) {
	fmt.Println(utils.SetRed("[SERVER]") + " Server active on:")
	if (listenAddr == "127.0.0.1" || listenAddr == "0.0.0.0") {
		fmt.Printf(utils.SetRed("[SERVER]") + "   - Internal: tcp://127.0.0.1:%d\n", port)
	}
	if (listenAddr != "127.0.0.1") {
		fmt.Printf(utils.SetRed("[SERVER]") + "   - External: tcp://%s:%d\n", listenAddr, port)
	}
	fmt.Println()
	fmt.Println(utils.SetGreen("[READY]") + "  Waiting for Client Handshakes...")
}