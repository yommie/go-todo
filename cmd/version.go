package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "Version",
	Short: "Print application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todos v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
