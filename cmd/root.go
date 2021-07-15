package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Initialize function will be called when every command gets called.
func init() {}

var rootCmd = &cobra.Command{
	Use:   "tpoller-server",
	Short: "Serve time-series data",
	Long:  `Serve time-series data from a connected Arduino device with an attached 'SparkFun Weather Shield' device over gRPC.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing...
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
