package internal

import (
	"log"
	"time"

	"google.golang.org/grpc"
	tstorage_pb "github.com/bartmika/tstorage-server/proto"

	pb "github.com/bartmika/tpoller-server/proto"
)

type TPoller struct {
	timer  *time.Timer
	ticker *time.Ticker
	done   chan bool

	readerFullAddress string
	readerConn              *grpc.ClientConn
	readerClient            pb.TPollerClient

	tstorageFullAddress string
	tstorageConn        *grpc.ClientConn
	tstorageClient      tstorage_pb.TStorageClient
}

func NewTPoller(
	readerFullAddress string,
	tstorageFullAddress string,
) (*TPoller, error) {
	s := &TPoller{
		timer:                   nil,
		ticker:                  nil,
		done:                    make(chan bool, 1), // Create a execution blocking channel.
		readerFullAddress:       readerFullAddress,
		tstorageFullAddress:     tstorageFullAddress,
	}

	// STEP 1: Connect to our time-series data storage.

	// Set up a direct connection to the gRPC server.
	conn, err := grpc.Dial(
		s.tstorageFullAddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	log.Println("tstorage connected")

	// Set up our protocol buffer interface.
	client := tstorage_pb.NewTStorageClient(conn)

	s.tstorageConn = conn
	s.tstorageClient = client

	// STEP 2: Connect to our serial reader.

	// Set up a direct connection to the gRPC server.
	conn, err = grpc.Dial(
		s.readerFullAddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	log.Println("treader connected")

	// Set up our protocol buffer interface.
	readerClient := pb.NewTPollerClient(conn)

	s.readerConn = conn
	s.readerClient = readerClient

	return s, nil
}

// Function will consume the main runtime loop and run the business logic
// of the application.
func (s *TPoller) RunMainRuntimeLoop() {
	defer s.shutdown()

	// DEVELOPERS NOTE:
	// (1) The purpose of this block of code is to find the future date where
	//     the minute just started, ex: 5:00 AM, 5:01, etc, and then start our
	//     main runtime loop to run along for every minute afterwords.
	// (2) If our application gets terminated by the user or system then we
	//     terminate our timer.
	log.Printf("Synching with local time...")
	s.timer = minuteTicker()
	select {
	case <-s.timer.C:
		log.Printf("Synchronized with local time.")
		s.ticker = time.NewTicker(1 * time.Minute)
	case <-s.done:
		s.timer.Stop()
		log.Printf("Interrupted timer.")
		return
	}

	// // THIS CODE IS FOR TESTING, REMOVE WHEN READY TO USE, UNCOMMENT ABOVE.
	// s.ticker = time.NewTicker(1 * time.Minute)

	// DEVELOPERS NOTE:
	// (1) The purpose of this block of code is to run as a goroutine in the
	//     background as an anonymous function waiting to get either the
	//     ticker chan or app termination chan response.
	// (2) Main runtime loop's execution is blocked by the `done` chan which
	//     can only be triggered when this application gets a termination signal
	//     from the operating system.
	log.Printf("TimeSeriesDataPoller is now running.")
	go func() {
		for {
			select {
			case <-s.ticker.C:
				err := s.pollArduinoReader()
				if err != nil {
					panic(err)
				}
			case <-s.done:
				s.ticker.Stop()
				log.Printf("Interrupted ticker.")
				return
			}
		}
	}()
	<-s.done
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (s *TPoller) StopMainRuntimeLoop() {
	s.done <- true
}

func (s *TPoller) shutdown() {
	s.tstorageConn.Close()
	s.readerConn.Close()
}
