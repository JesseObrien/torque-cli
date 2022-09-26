package run

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	RunCmd = &cobra.Command{
		Use:   "run",
		Short: "Run a Torque app.",
		Long:  "",
		Run:   executeRun,
	}
)

func executeRun(cmd *cobra.Command, args []string) {
	log.Info(fmt.Sprintf("🔨 running %s", viper.GetViper().GetString("app.name")))

	rc := exec.CommandContext(cmd.Context(), "go", "run", "cmd/main/main.go")
	rc.Stdout = os.Stdout
	rc.Stderr = os.Stderr

	out, err := rc.Output()
	if err != nil {
		log.Errorf("Failed to run `go run`: %s", err.Error())
	}

	log.Info(string(out))
}
