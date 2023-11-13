package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/useloopso/BMR/config"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "bmr",
	Version: version,
	Short:   "Entrypoint for the BMR node",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := config.Load(); err != nil {
			fmt.Println("Failed to load config: ", err)
			os.Exit(1)
		}

		if err := config.LoadChainInfo(); err != nil {
			fmt.Println("Failed to load chain info: ", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed to start BMR: ", err)
		os.Exit(1)
	}
}
