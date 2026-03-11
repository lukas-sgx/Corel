package server

import (
	"fmt"
)

var header = 
ClearScreen + Blue +
"==============================\n" + Bold +
"[ COREL  :: AUTONOMOUS AGENT ]\n" + Reset + Blue +
"[ MODE   ::           SERVER ]\n" +
"[ STATUS ::        LISTENING ]\n" +
"==============================\n" + Reset

func initHeader() {
	fmt.Println(header)
	fmt.Println(setBlue("[SCAN]") + " Listening for incoming beacons...")
}

func isMeshSync(mesh bool) {
	if (mesh == true) {
		fmt.Println(setBlue("[SCAN]") + " Global mesh synchronization: ACTIVE signification\n");
	} else {
		fmt.Println()
	}
}

func listenerInfo(listenAddr string, port int) {
	fmt.Println(setRed("[SERVER]") + " Server active on:")
	if (listenAddr == "127.0.0.1" || listenAddr == "0.0.0.0") {
		fmt.Printf(setRed("[SERVER]") + "   - Internal: tcp://127.0.0.1:%d\n", port)
	}
	if (listenAddr != "127.0.0.1") {
		fmt.Printf(setRed("[SERVER]") + "   - External: tcp://%s:%d\n", listenAddr, port)
	}
	fmt.Println()
	fmt.Println(setGreen("[READY]") + "  Waiting for Client Handshakes...")
}