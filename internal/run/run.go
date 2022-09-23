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
		Use:   "watch",
		Short: "Watch and re-run a Torque app. Use --build to output a binary as well.",
		Long:  "",
		Run:   executeRun,
	}
)

func executeRun(cmd *cobra.Command, args []string) {
	log.Info(fmt.Sprintf("🔨 running %s", viper.GetViper().GetString("app.name")))

	rc := exec.Command("go", "run", "cmd/main/main.go")
	rc.Stdout = os.Stdout
	rc.Stderr = os.Stderr

	if err := rc.Run(); err != nil {
		log.Errorf("Failed to run `go run`: %s", err.Error())
	}
}
