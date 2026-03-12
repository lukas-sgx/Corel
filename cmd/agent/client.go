package agent

import (
	"github.com/spf13/cobra"
)

type AgentInfo struct {
	remoteAddr string
	remotePort int
	meshSync bool
	version string
	identity string
}

var (
	remoteAddr string
	remotePort int
	meshSync bool
	version string
)

var ClientCmd = &cobra.Command{
	Use:     "client",
	Short:   "Start the Corel client",
	Long:    `Start the Corel client, it will connect to the specified server and handle incoming requests.`,
	GroupID: "modules",
	Run: clientConnect,
}

func clientConnect(cmd *cobra.Command, args []string) {

	agent := AgentInfo{
		remoteAddr: remoteAddr,
		remotePort: remotePort,
		meshSync: meshSync,
		identity: generateIdentity(),
		version: version,
	}
	initHeader(agent)
}

func InitFlags(versionCli string) {
	version = versionCli
	ClientCmd.Flags().StringVarP(&remoteAddr, "RHOST", "i", "127.0.0.1", "Ip address server")
	ClientCmd.Flags().IntVarP(&remotePort, "RPORT", "p", 8080, "Ip port server")
	ClientCmd.Flags().BoolVarP(&meshSync, "MSYNC", "m", true, "Use mesh sync")
}
