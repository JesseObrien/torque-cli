package config

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"
)

var (
	// Whether to make a local folder .torque.yml file or not
	local bool

	CfgCmd = &cobra.Command{
		Use:   "config",
		Short: "Modify the default configuration for torque-cli.",
		Long:  "",
		Run:   executeConfig,
	}
)

func init() {
	CfgCmd.PersistentFlags().BoolVar(&local, "local", false, "If local is true, torque will set a config value in the .torque.yml file in the current directory.")
	log.SetHandler(cli.New(os.Stderr))
}

func executeConfig(cmd *cobra.Command, args []string) {
}
