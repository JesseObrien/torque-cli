package start

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jesseobrien/torque/internal/scaffold"
	"github.com/spf13/cobra"
)

var (
	dry        bool
	moduleName string

	InitCmd = &cobra.Command{
		Use:   "init [appname]",
		Short: "initialize a new Torque project directory",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run:   executeInit,
	}
)

func init() {
	InitCmd.PersistentFlags().BoolVar(&dry, "dry-run", false, "Whether torque will do a dry run of scaffolding everything and clean up after.")
	InitCmd.PersistentFlags().StringVar(&moduleName, "mod-name", "", "The go module name that will be used to initialize go.mod. If none is specified, the project name is used.")
	log.SetHandler(cli.New(os.Stderr))
}

func executeInit(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Error("Required `appname` arg is missing.")
		return
	}

	appName := args[0]
	if !isValidAppName(appName) {
		log.Error("`appname` argument is invalid. App names can only contain letters, numbers and - or _.")
		return
	}

	log.Info(fmt.Sprintf("🔨 Creating new app directory: %s", appName))

	if err := createProjectDirectories(appName); err != nil {
		cleanupProjectDirectory(appName)
		log.WithError(err).Error("creating root directory failed")
		return
	}

	cfg := scaffold.ScaffoldConfig{AppName: appName}

	s := scaffold.NewScaffolder(cfg)

	if err := s.Scaffold(); err != nil {
		cleanupProjectDirectory(appName)
		log.WithError(err).Error("scaffolding project files failed")
		return
	}

	log.Info("Done. Happy building!")

	if dry {
		log.Info("Dry-run is enabled. Cleaning up...")
		os.Remove(appName)
	}

}

func cleanupProjectDirectory(appName string) {
	os.Remove(appName)
}

func createProjectDirectories(appName string) error {
	exists, err := os.Stat(appName)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if exists != nil {
		files, err := os.ReadDir(appName)
		if err != nil {
			return err
		}

		if len(files) > 0 {
			return fmt.Errorf("'%s' directory exists and is not empty", appName)
		}
	}

	if _, err := os.Stat(appName); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(appName, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if err := os.Chdir(appName); err != nil {
		return err
	}

	directoryTree := []string{
		"cmd/main",
		"internal/http",
		"internal/data",
		"tmp",
	}

	for _, directory := range directoryTree {
		if err := os.MkdirAll(directory, os.ModePerm); err != nil {
			cleanupProjectDirectory(appName)

			return err
		}
	}

	return nil
}

// Validate that the app name will work as a directory name

func isValidAppName(name string) bool {
	reg := regexp.MustCompile(`^[^\\/()!?%*:|"<>\.]+$`)

	return reg.MatchString(name)
}
