package internal

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jesseobrien/torque-cli/internal/config"
	"github.com/jesseobrien/torque-cli/internal/new"
	"github.com/jesseobrien/torque-cli/internal/run"
	"github.com/jesseobrien/torque-cli/internal/watch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Configuration file path
	cfgFilePath string

	rootCmd = &cobra.Command{
		Use:   "torque",
		Short: "CLI interface for building powerful web apps.",
		Long:  `Torque as a web framework uses this CLI to empower building web applications.`,
	}
)

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error(err.Error())
	}

}

func init() {
	log.SetHandler(cli.New(os.Stderr))

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "", "config file path (default is $PWD/torque.yml)")

	rootCmd.AddCommand(new.InitCmd)
	rootCmd.AddCommand(config.CfgCmd)
	rootCmd.AddCommand(watch.WatchCmd)
	rootCmd.AddCommand(run.RunCmd)
}

func initConfig() {

	// Get the current working directory path

	viper.AutomaticEnv()

	path, err := os.Getwd()
	cobra.CheckErr(err)

	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	viper.SetConfigName("torque.yml")

	err = viper.ReadInConfig()
	if err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}

	// if err != nil {
	// 	log.Errorf("Error loading config: %s", err)
	// }

	// @TODO: Mayyyybe use this. It'll be for a global config file we might not need

	// if cfgFilePath != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFilePath)
	// } else {
	// 	// Search config in home directory with name ".cobra" (without extension).
	// 	viper.AddConfigPath(path)
	// 	viper.SetConfigType("yml")
	// 	viper.SetConfigName("torque.yml")
	// }

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	log.Infof("Using config file:", viper.ConfigFileUsed())
	// }

	// if err != nil {
	// 	log.Errorf("Error loading config: %s", err)
	// }
}
