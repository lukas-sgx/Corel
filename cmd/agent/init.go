package agent

import (
	"github.com/lukas-sgx/corel/internal/agent"
	"github.com/lukas-sgx/corel/pkg/utils"

	"fmt"
	"math/rand"
	"os"
)

func generateIdentity() string {
    uid, _ := os.Hostname()
    id := rand.Uint32()
    
    return fmt.Sprintf("%s-%d", uid, id)
}

func initHeader(agent agent.Agent) {
	var header =  utils.ClearScreen + utils.Red +
`     .
    / \      ` + utils.Blue + utils.Bold + `[ COREL :: AUTONOMOUS AGENT ]` + utils.Reset + utils.Red + `
   /   \     ` + utils.Blue + `[ VERSION : `+ agent.Version +`           ]` + utils.Red + `
  /_____\    ` + utils.Blue + `[ STATUS  : INITIALIZING    ]` + utils.Red + `
  \     /
   \   /     ` + utils.Blue + `TARGET:   tcp://`+ agent.RemoteAddr +`:`+ fmt.Sprintf("%d", agent.RemotePort) + utils.Red + `
    \ /      ` + utils.Blue + `IDENTITY: `+ agent.Identity + utils.Red + `
     '` + utils.Reset

	fmt.Println(header)
}