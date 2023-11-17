package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
			return
		}

		stopCh := make(chan struct{})
		go loopsoClient.ListenAll(stopCh)

		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

		<-signalCh

		close(stopCh)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
