# TPoller Server

The purpose of this application is to poll time-series data per time interval from any application which implements the [the following gRPC function](/proto/tpoller.proto):

```proto
service TPoller {
    rpc PollTimeSeriesData (google.protobuf.Empty) returns (stream PolledTimeSeriesDatum) {}
}

message PolledLabel {
    string name = 1;
    string value = 2;
}

message PolledTimeSeriesDatum {
    string metric = 1;
    repeated PolledLabel labels = 2;
    double value = 3;
    google.protobuf.Timestamp timestamp = 4;
}
```

The interval of time is every one minute. For example, if your application implemented this gRPC protocol and is running, this poller will request time-series data for following times will result in a poll.
- 2021/07/13 23:17:00
- 2021/07/13 23:18:00
- 2021/07/13 23:19:00
- 2021/07/13 23:20:00
- ... etc

If you'd like to see how to setup a server, see [tarduinoreader-server](https://github.com/bartmika/tarduinoreader-server).

## Prerequisites

You must have the following installed before proceeding. If you are missing any one of these then you cannot begin.

* ``Go 1.16.3``

## Installation
<!-- 1. Please visit the [sparkfunweathershield-arduino](https://github.com/bartmika/sparkfunweathershield-arduino) repository and setup the external device and connect it to your development machine.

2. Please visit the [serialreader-server](https://github.com/bartmika/serialreader-server) repository and setup that application on your device.

3. Please visit the [tstorage-server](https://github.com/bartmika/tstorage-server) repository and setup that application on your device. -->

4. Get our latest code.

    ```bash
    go get -u github.com/bartmika/tpoller-server
    ```

5. Setup our environment variable before running our server.

    ```
    export POLLER_SERVER_SERIAL_READER_SERVER_ADDRESS=127.0.0.1
    export POLLER_SERVER_SERIAL_READER_SERVER_PORT=50052
    export POLLER_SERVER_TSTORAGE_SERVER_ADDRESS=127.0.0.1
    export POLLER_SERVER_TSTORAGE_SERVER_PORT=50051
    ```

## Usage
Run our application.

    go run main.go serve

If the server successfully starts you should see a message in your **termnal**:

    2021/07/10 15:40:36 Synching with local time...
    2021/07/10 15:41:00 Synchronized with local time.
    2021/07/10 15:41:00 TPoller is now running.

## License

This application is licensed under the **BSD 3-Clause License**. See [LICENSE](LICENSE) for more information.
