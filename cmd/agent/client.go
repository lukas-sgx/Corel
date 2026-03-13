package agent

import (
	"github.com/spf13/cobra"
	"github.com/lukas-sgx/corel/internal/agent"
)

var (
	remoteAddr string
	remotePort int
	meshSync bool
	version string
	secrets bool
)

var ClientCmd = &cobra.Command{
	Use:     "client",
	Short:   "Start the Corel client",
	Long:    `Start the Corel client, it will connect to the specified server and handle incoming requests.`,
	GroupID: "modules",
	Run: clientConnect,
}

func clientConnect(cmd *cobra.Command, args []string) {

	agentInfo := agent.Agent{
		RemoteAddr: remoteAddr,
		RemotePort: remotePort,
		MeshSync: meshSync,
		Version: version,
		Identity: generateIdentity(),
		Secrets: secrets,
	}

	if secrets == false {
		initHeader(agentInfo)
	}
	agent.Peer(agentInfo)
}

func InitFlags(versionCli string) {
	version = versionCli
	ClientCmd.Flags().StringVarP(&remoteAddr, "RHOST", "i", "127.0.0.1", "Ip address server")
	ClientCmd.Flags().IntVarP(&remotePort, "RPORT", "p", 8080, "Ip port server")
	ClientCmd.Flags().BoolVarP(&meshSync, "MSYNC", "m", true, "Use mesh sync")
	ClientCmd.Flags().BoolVarP(&secrets, "SECRETS", "s", false, "Use secrets")
}
