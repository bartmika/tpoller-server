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
	Short: "Poll time-series data and save to storage",
	Long:  `Connect to a 'tstorage-server' and a 'data reader' to poll time-series data and save it to the storage server.`,
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
