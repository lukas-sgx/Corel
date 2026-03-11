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

func InitHeader() {
	fmt.Println(header)
}