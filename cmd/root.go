package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

type Config struct {
	maxINSERTS int
	minINSERTS int
	maxUPDATES int
	minUPDATES int
}

var cfg Config

var rootCmd = &cobra.Command{
	Use:   "psql-stress-test",
	Short: "A simple CLI tool to stress test postgress",
	Long: `A simple CLI tool to stress test postgress`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.psql-stress-test.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".psql-stress-test" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("config")
		viper.SetConfigName(".psql-stress-test")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		cfg.maxINSERTS = viper.GetInt("Config.maxINSERTS")
		cfg.minINSERTS = viper.GetInt("Config.minINSERTS")
		cfg.maxUPDATES = viper.GetInt("Config.maxUPDATES")
		cfg.minUPDATES = viper.GetInt("Config.minUPDATES")
	} else {
		fmt.Println("Config file not found, using default values")
		cfg.maxINSERTS = 10
		cfg.minINSERTS = 10
		cfg.maxUPDATES = 10
		cfg.minUPDATES = 10
	}
}
