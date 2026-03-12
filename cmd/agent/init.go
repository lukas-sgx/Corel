package agent

import (
	"github.com/lukas-sgx/corel/pkg/utils"
	"fmt"
    "os"
    "math/rand"
)

func generateIdentity() string {
    uid, _ := os.Hostname()
    id := rand.Uint32()
    
    return fmt.Sprintf("%s-%d", uid, id)
}

func initHeader(agent AgentInfo) {
	var header =  utils.ClearScreen + utils.Red +
`     .
    / \      ` + utils.Blue + utils.Bold + `[ COREL :: AUTONOMOUS AGENT ]` + utils.Reset + utils.Red + `
   /   \     ` + utils.Blue + `[ VERSION : `+ agent.version +`           ]` + utils.Red + `
  /_____\    ` + utils.Blue + `[ STATUS  : INITIALIZING    ]` + utils.Red + `
  \     /
   \   /     ` + utils.Blue + `TARGET:   tcp://`+ agent.remoteAddr +`:`+ fmt.Sprintf("%d", agent.remotePort) + utils.Red + `
    \ /      ` + utils.Blue + `IDENTITY: `+ agent.identity + utils.Red + `
     '` + utils.Reset

	fmt.Println(header)
}