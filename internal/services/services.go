package services

import (
	"context"
	"fmt"

	"os/exec"

	"github.com/apex/log"
	"github.com/spf13/cobra"
)

var (
	RunCmd = &cobra.Command{
		Use:   "services [command]",
		Short: "Run docker-compose",
		Run:   executeRun,
	}
)

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
	cmd.Dir = "/Users/keith/tmp/myApp"

	err := cmd.Start()
	if err != nil {
		log.Errorf("Failed to run `docker-compose up`: %s", err.Error())
	}
	log.Infof("Waiting for command to finish...")
	err = cmd.Wait()
	log.Errorf("Command finished with error: %v", err)
}
