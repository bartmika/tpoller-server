package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	serialReaderAddress string
	serialReaderPort string
	tstorageAddress string
	tstoragePort string
)

// Initialize function will be called when every command gets called.
func init() {
	// Load up our `environment variables` from our operating system.
	srAddress := os.Getenv("POLLER_SERVER_SERIAL_READER_SERVER_ADDRESS")
	srPort := os.Getenv("POLLER_SERVER_SERIAL_READER_SERVER_PORT")
	tsAddress := os.Getenv("POLLER_SERVER_TSTORAGE_SERVER_ADDRESS")
	tsPort := os.Getenv("POLLER_SERVER_TSTORAGE_SERVER_PORT")

	// Make our environment variable globally accessible throughout all the
	// sub-commands in this project.
	rootCmd.PersistentFlags().StringVar(&serialReaderAddress, "srAddress", srAddress, "The address of the serial reader server this application will connect to.")
	rootCmd.PersistentFlags().StringVar(&serialReaderPort, "srPort", srPort, "The port of the serial reader server this application will connect to.")
	rootCmd.PersistentFlags().StringVar(&tstorageAddress, "tsAddress", tsAddress, "The address of the tstorage server this application will connect to.")
	rootCmd.PersistentFlags().StringVar(&tstoragePort, "tsPort", tsPort, "The port of the tstorage server this application will connect to.")
}

var rootCmd = &cobra.Command{
	Use:   "poller-server",
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
