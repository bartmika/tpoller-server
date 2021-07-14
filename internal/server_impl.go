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

	// pb "github.com/bartmika/tpoller-server/proto"
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

// func (s *TPoller) getDataFromArduino() *pb.SparkFunWeatherShieldTimeSeriesData {
// 	// c := s.readerClient
// 	//
// 	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	// defer cancel()
// 	// r, err := c.GetSparkFunWeatherShieldData(ctx, &pb.GetTimeSeriesData{})
// 	// if err != nil {
// 	// 	log.Fatalf("could not greet: %v", err)
// 	// }
// 	// return r
// }

func (s *TPoller) pollArduinoReader() error {
	c := s.readerClient

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Perform our gRPC request.
	pollStream, err := c.PollTimeSeriesData(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not select: %v", err)
	}

	// Handle our pollStream of data from the server.
	for {
		timeSeriesData, err := pollStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error with pollStream: %v", err)
		}

		// Print out the gRPC response.
		log.Printf("Server Response: %s", timeSeriesData)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.GetSparkFunWeatherShieldData(ctx, &pb.GetTimeSeriesData{})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// return r
	return nil
}

// func (s *TPoller) saveDataToStorage(data *pb.SparkFunWeatherShieldTimeSeriesData) {
// 	// // For debugging purposes only.
// 	// // fmt.Printf("\n%+v\n", data)
// 	//
// 	// // Open up a streamming service connection with our `tstorage-server` so
// 	// // we can send bulk time-series data and gRPC will send the data in streams.
// 	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	// defer cancel()
// 	// stream, err := s.tstorageClient.InsertRows(ctx)
// 	// if err != nil {
// 	// 	log.Fatalf("%v.InsertRows(_) = _, %v", s.tstorageClient, err)
// 	// }
// 	//
// 	// // Send the data.
// 	// s.addTimeSeriesDatum(stream, TPollerSparkFunWeatherShieldHumiditySensorId, data.HumidityValue, data.Timestamp)
// 	// s.addTimeSeriesDatum(stream, TPollerSparkFunWeatherShieldTemperatureSensorId, data.TemperatureValue, data.Timestamp)
// 	// s.addTimeSeriesDatum(stream, TPollerSparkFunWeatherShieldPressureSensorId, data.PressureValue, data.Timestamp)
// 	// s.addTimeSeriesDatum(stream, TPollerSparkFunWeatherShieldTemperatureBackupSensorId, data.TemperatureBackupValue, data.Timestamp)
// 	// s.addTimeSeriesDatum(stream, TPollerSparkFunWeatherShieldAltitudeSensorId, data.AltitudeValue, data.Timestamp)
// 	// s.addTimeSeriesDatum(stream, TPollerSparkFunWeatherShieldIlluminanceSensorId, data.IlluminanceValue, data.Timestamp)
// 	//
// 	// // Terminate our streamming connection. Ignore the server reply message sent.
// 	// _, err = stream.CloseAndRecv()
// 	// if err != nil {
// 	// 	log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
// 	// }
// 	// // log.Printf("Server Response: %v", reply)
// }
//
// func (s *TPoller) addTimeSeriesDatum(
// // 	stream tstorage_pb.TStorage_InsertRowsClient,
// // 	instrumentId string,
// // 	value32 float32,
// // 	ts *timestamp.Timestamp,
// // ) {
// // 	// Generate our labels.
// // 	labels := []*tstorage_pb.Label{}
// // 	labels = append(labels, &tstorage_pb.Label{Name: "Source", Value: "Command"})
// //
// // 	// DEVELOPERS NOTE:
// // 	// The hardware returns a `float32` value but our database stores in `float64`
// // 	// so as a result we will cast into the database prefered format.
// // 	value64 := float64(value32)
// //
// // 	tsd := &tstorage_pb.TimeSeriesDatum{
// // 		Labels:    labels,
// // 		Metric:    instrumentId,
// // 		Value:     value64,
// // 		Timestamp: ts,
// // 	}
// //
// // 	// DEVELOPERS NOTE:
// // 	// To stream from a client to a server using gRPC, the following documentation
// // 	// will help explain how it works. Please visit it if the code below does
// // 	// not make any sense.
// // 	// https://grpc.io/docs/languages/go/basics/#client-side-streaming-rpc-1
// //
// // 	if err := stream.Send(tsd); err != nil {
// // 		log.Fatalf("could not add time-series data to storage: %v", err)
// // 	}
// }
