package internal

import (
	// "fmt"
	"context"
	"log"
	"time"
	"io"

	// "google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	// "github.com/golang/protobuf/ptypes/timestamp"
	// tstorage_pb "github.com/bartmika/tstorage-server/proto"

	pb "github.com/bartmika/tpoller-server/proto"
)

// Source: https://www.reddit.com/r/golang/comments/44tmti/scheduling_a_function_call_to_the_exact_start_of/
func minuteTicker() *time.Timer {
	// Current time
	now := time.Now()

	// Get the number of seconds until the next minute
	var d time.Duration
	d = time.Second * time.Duration(60-now.Second())

	// Time of the next tick
	nextTick := now.Add(d)

	// Subtract next tick from now
	diff := nextTick.Sub(time.Now())

	// Return new ticker
	return time.NewTimer(diff)
}

func (s *TPoller) pollArduinoReader() error {
	c := s.telemeterClient

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Open up a streaming service connection with our application that implemented
	// our gRPC definition.
	telemetryStream, err := c.PollTelemeter(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not select: %v", err)
	}

	// Open up a streamming service connection with our `tstorage-server` so
	// we can send bulk time-series data and gRPC will send the data in streams.
	storageStream, err := s.tstorageClient.InsertRows(ctx)
	if err != nil {
		log.Fatalf("%v.InsertRows(_) = _, %v", s.tstorageClient, err)
	}

	// Handle our telemetryStream of data from the server.
	for {
		telemetryDatum, err := telemetryStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error with telemetryStream: %v", err)
		}

		// Print out the gRPC response.
		log.Printf("Server Response: %s", telemetryDatum)

		// Convert from our `polled` format to the storage format.
		timeSeriesDatum := pb.ToTimeSeriesDatum(telemetryDatum)

		// DEVELOPERS NOTE:
		// To stream from a client to a server using gRPC, the following documentation
		// will help explain how it works. Please visit it if the code below does
		// not make any sense.
		// https://grpc.io/docs/languages/go/basics/#client-side-streaming-rpc-1

		if err := storageStream.Send(timeSeriesDatum); err != nil {
			log.Fatalf("could not add time-series data to storage: %v", err)
		}
	}

	_, err = storageStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", storageStream, err, nil)
	}
	// log.Printf("Server Response: %v", reply)

	return nil
}
