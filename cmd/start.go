package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/useloopso/BMR/config"
	"github.com/useloopso/BMR/core"
)

var startCmd = &cobra.Command{
	Use:  "start",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		loopsoClient, err := core.NewClient(config.Chains)
		if err != nil {
			fmt.Println("Failed to init loopso client: ", err)
		}

		loopsoClient.ListenAll()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
