package cmd

import (
	"log"

	"github.com/spf13/viper"
	"github.com/yommie/go-todo/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "Todo Web Application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(utils.GetProjectRoot() + "/configs")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

// Execute commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
