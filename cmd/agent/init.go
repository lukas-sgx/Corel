package agent

import (
	"github.com/lukas-sgx/corel/pkg/utils"
	"fmt"
)

func InitHeader(agent AgentInfo) {
	var header =  utils.ClearScreen + utils.Red +
`     .
    / \      ` + utils.Blue + `[ COREL :: AUTONOMOUS AGENT ]` + utils.Red + `
   /   \     ` + utils.Blue + `[ VERSION : `+ agent.version +`           ]` + utils.Red + `
  /_____\    ` + utils.Blue + `[ STATUS  : INITIALIZING    ]` + utils.Red + `
  \     /
   \   /     ` + utils.Blue + `TARGET:   tcp://`+ agent.remoteAddr +`:`+ fmt.Sprintf("%d", agent.remotePort) + utils.Red + `
    \ /      ` + utils.Blue + `IDENTITY: `+ agent.identity + utils.Red + `
     '` + utils.Reset

	fmt.Println(header)
}