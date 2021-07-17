# TPoller Server
## Overview
The purpose of this application is to poll time-series data per time interval from any (telemetry) application running a gRPC server which implements the [the following gRPC service definition](/proto/telemetry.proto):

```protobuf
service Telemetry {
    rpc GetTimeSeriesData (google.protobuf.Empty) returns (stream TelemetryDatum) {}
}

message TelemetryLabel {
    string name = 1;
    string value = 2;
}

message TelemetryDatum {
    string metric = 1;
    repeated TelemetryLabel labels = 2;
    double value = 3;
    google.protobuf.Timestamp timestamp = 4;
}
```

Your server which implemented the gRPC service definition is called a **data reader**. `tpoller-server` will then send your polled data over gRPC to a **fast time-series data storage server** called [tstorage-server](https://github.com/bartmika/tstorage-server).

## Installation

Get our latest code.

```bash
go install github.com/bartmika/tpoller-server@latest
```

## Usage
Run our application.

```bash
$GOBIN/tpoller-server serve --telemetry_addr="127.0.0.1:50051" --storage_addr="127.0.0.1:50052"
```

If the server successfully starts you should see a message in your **termnal**:

```bash
2021/07/15 00:51:00 Storage connected
2021/07/15 00:51:00 Telemeter connected
2021/07/15 00:51:00 Synching with local time...
2021/07/15 00:52:00 Synchronized with local time.
2021/07/15 00:52:00 TPoller is now running.
```

More sub-command details:


```text
Run the tpoller service in the foreground which will periodically call the "data reader" to retrieve time-series data and save it to "tstorage" server.

Usage:
  tpoller-server serve [flags]

Flags:
  -h, --help                    help for serve
  -o, --storage_addr string     The time-series data storage gRPC server address. (default "localhost:50051")
  -i, --telemetry_addr string   The telemetry gRPC server address with the 'data reader'. (default "localhost:50052")
```

## Used by:
This server is confirmed to successfully poll from the following application(s):
* [treader-server](https://github.com/bartmika/treader-server)

If you'd like to add your app, please create a [pull request](https://github.com/bartmika/tpoller-server/pulls) with a link to your app.

## Contributing
### Development
If you'd like to setup the project for development. Here are the installation steps:

1. Go to your development folder.

    ```bash
    cd ~/go/src/github.com/bartmika
    ```

2. Clone the repository.

    ```bash
    git clone https://github.com/bartmika/tpoller-server.git
    cd tpoller-server
    ```

3. Install the package dependencies

    ```bash
    go mod tidy
    ```

4. In your **terminal**, make sure we export our path (if you haven’t done this before) by writing the following:

    ```bash
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

5. Run the following to generate our new gRPC interface. Please note in your development, if you make any changes to the gRPC service definition then you'll need to rerun the following:

    ```bash
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/telemetry.proto
    ```

6. You are now ready to start the server and begin contributing!

    ```bash
    go run main.go serve --telemetry_addr="127.0.0.1:50051" --storage_addr="127.0.0.1:50052"
    ```

### Quality Assurance

Found a bug? Need Help? Please create an [issue](https://github.com/bartmika/tpoller-server/issues).


## License

This application is licensed under the **BSD 3-Clause License**. See [LICENSE](LICENSE) for more information.
