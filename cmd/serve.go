package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/bartmika/tpoller-server/internal"
)

var (
	telemetryAddr string
	storageAddr   string
)

func init() {
	// The following are required.
	serveCmd.Flags().StringVarP(&telemetryAddr, "telemetry_addr", "i", "localhost:50052", "The telemetry gRPC server address with the 'data reader'.")
	serveCmd.MarkFlagRequired("telemetry_addr")
	serveCmd.Flags().StringVarP(&storageAddr, "storage_addr", "o", "localhost:50051", "The time-series data storage gRPC server address.")
	serveCmd.MarkFlagRequired("storage_addr")

	// Attach our sub-command to our application
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the tpoller service",
	Long:  `Run the tpoller service in the foreground which will periodically call the "data reader" to retrieve time-series data and save it to "tstorage" server.`,
	Run: func(cmd *cobra.Command, args []string) {
		doRun()
	},
}

func doRun() {
	app, err := internal.NewTPoller(telemetryAddr, storageAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer app.StopMainRuntimeLoop()

	// DEVELOPERS CODE:
	// The following code will create an anonymous goroutine which will have a
	// blocking chan `sigs`. This blocking chan will only unblock when the
	// golang app receives a termination command; therfore the anyomous
	// goroutine will run and terminate our running application.
	//
	// Special Thanks:
	// (1) https://gobyexample.com/signals
	// (2) https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	//
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs // Block execution until signal from terminal gets triggered here.
		fmt.Println("Starting graceful shut down now.")
		app.StopMainRuntimeLoop()
	}()

	app.RunMainRuntimeLoop()
}
