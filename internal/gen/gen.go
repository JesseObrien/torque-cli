package gen

import (
	"github.com/spf13/cobra"
)

var (
	GenCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate something useful inside of a Torque app.",
		Long:  "",
	}
)

func init() {
	GenCmd.AddCommand(ControllerCmd)
}
