package internal

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/jesseobrien/torque-cli/internal/config"
	"github.com/jesseobrien/torque-cli/internal/new"
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "", "config file path (default is $PWD/torque.yml)")

	rootCmd.AddCommand(new.InitCmd)
	rootCmd.AddCommand(config.CfgCmd)
}

func initConfig() {
	if cfgFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFilePath)
	} else {
		// Find home directory.
		path, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path)
		viper.SetConfigType("yml")
		viper.SetConfigName("torque")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
