package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"os/exec"

	"github.com/apex/log"
	"github.com/spf13/cobra"
)

var (
	workingDir string
	InitCmd    = &cobra.Command{
		Use:   "services [command]",
		Short: "Run docker-compose",
		Run:   executeRun,
	}
)

func init() {
	InitCmd.PersistentFlags().StringVar(&workingDir, "working-dir", "", "Provide a working directory for command.")
}
func executeRun(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "up":
		log.Info(fmt.Sprintf("🔨 running docker-compose %s", args[0]))
		dockerComposeUp(cmd.Context())
	case "down":
		log.Info(fmt.Sprintf("🔨 running docker-compose %s", args[0]))
	}
}
func dockerComposeUp(ctx context.Context) {
	//TODO: Get Values from config with viper and create COMPOSE_PROFILES

	cmd := exec.CommandContext(ctx, "docker-compose", "up")
	cmd.Env = append(cmd.Env, "COMPOSE_PROFILES=redis")
	//TODO: Will add dir flag to set WD if not current dir
	if workingDir != "" {
		cmd.Dir = workingDir
	} else {
		wd, err := os.Getwd()
		if err != nil {
			log.Errorf("Failed to get working dir: %s", err.Error())
		}
		workingDir = wd
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Errorf("Failed to run `docker-compose up`: %s", err.Error())
	}

	log.Infof("Waiting for command to finish...")

	err = cmd.Wait()
	log.Errorf("Command finished with error: %v", err)

	// Inspect the error a bit further to give a recommended action
	if strings.Contains(err.Error(), "docker-compose-v1") {
		log.Errorf("ℹ️ docker-compose might not be installed. Try installing docker-compose")
	}
}
