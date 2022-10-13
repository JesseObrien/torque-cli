package new

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/apex/log"
	"github.com/jesseobrien/torque-cli/internal/scaffold"
	"github.com/spf13/cobra"
)

var (
	dry        bool
	moduleName string
	customPath string
	orm        bool
	redis      bool

	InitCmd = &cobra.Command{
		Use:   "new [appname]",
		Short: "Initialize a new Torque project directory",
		Long:  "",
		Args:  cobra.MinimumNArgs(1),
		Run:   executeInit,
	}
)

func init() {
	InitCmd.PersistentFlags().BoolVar(&dry, "dry-run", false, "Whether torque will do a dry run of scaffolding everything and clean up after.")
	InitCmd.PersistentFlags().BoolVar(&orm, "orm", true, "If orm is false, torque will not generate ORM database files.")
	InitCmd.PersistentFlags().BoolVar(&redis, "redis", false, "If redis is true, torque will enable it in docker-compose.")
	InitCmd.PersistentFlags().StringVar(&customPath, "dir", "", "Provide a custom path to scaffold the project into.")
	InitCmd.PersistentFlags().StringVar(&moduleName, "mod-name", "", "The go module name that will be used to initialize go.mod. If none is specified, the project name is used.")
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

	appDir := "./"
	if customPath != "" {
		if strings.HasSuffix(customPath, "/") {
			customPath = customPath[:len(customPath)-len("/")]
		}
		appDir = customPath
	}

	appDir = fmt.Sprintf("%s/%s", appDir, appName)

	log.Infof("🔨 Creating new app in directory %s", appDir)

	if err := createProjectDirectories(appDir); err != nil {
		cleanupProjectDirectory(appDir)
		log.WithError(err).Error("creating root directory failed")
		return
	}

	modName := appName
	if moduleName != "" {
		modName = moduleName
	}

	cfg := scaffold.ScaffoldConfig{
		AppName: appName,
		ModName: modName,
		Path:    appDir,
		ORM:     orm,
		Redis:   redis,
	}

	s := scaffold.NewScaffolder(cfg)

	if err := s.ScaffoldNewProject(); err != nil {
		cleanupProjectDirectory(appDir)
		log.WithError(err).Error("scaffolding project files failed")
		return
	}

	if err := initializeGoModule(modName); err != nil {
		log.WithError(err).Error("go mod init failed")
		return
	}

	log.Info("✅ Done. Happy building!")

	if dry {
		log.Info("Dry-run is enabled. Cleaning up...")
		os.Remove(appDir)
	}

}

func cleanupProjectDirectory(appName string) {
	os.Remove(appName)
}

func initializeGoModule(modName string) error {
	log.Infof("🔨 Running `go mod init %s`", modName)

	cmd := exec.Command("go", "mod", "init", modName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	log.Infof("🔨 Running `go mod tidy`")

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func createProjectDirectories(appName string) error {
	log.Info("🔨 Scaffolding project directories.")
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
		err := os.MkdirAll(appName, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if err := os.Chdir(appName); err != nil {
		return err
	}

	directoryTree := []string{
		"cmd/main",
		"dist",
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
