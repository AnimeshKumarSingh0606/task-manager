package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/Gophercizes/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A simple CLI to remember TODOs",
	Long:  `When you has something to do, just use task to add it to the TODO list`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".task")
	}

	viper.AutomaticEnv() 

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
