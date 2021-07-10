package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"

	"github.com/spf13/cobra"

	"github.com/bartmika/poller-server/internal"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the poller service",
	Long:  `Run the poller service in the foregone which will periodically call the "serialreader-server" to retrieve time-series data and save it to our database.`,
	Run: func(cmd *cobra.Command, args []string) {
		doRun()
	},
}

func doRun() {
	serialReaderFullAddress := fmt.Sprintf("%v:%v", serialReaderAddress, serialReaderPort)
	tstorageFullAddress := fmt.Sprintf("%v:%v", tstorageAddress, tstoragePort)

	app, err := internal.NewPollerServer(serialReaderFullAddress, tstorageFullAddress)
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
