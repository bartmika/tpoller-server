```
cd ~/go/src/github.com/bartmika
git clone https://github.com/bartmika/tpoller-server.git
cd tpoller-server
```


Initialize golang modules.

```
go mod init github.com/bartmika/tstorage-server
```

Install our project’s dependencies.

```
export GO111MODULE=on  # Enable module mode
go mod tidy
```

In your terminal, make sure we export our path (if you haven’t done this before) by writing the following:

```
export PATH="$PATH:$(go env GOPATH)/bin"
```

Run the following to generate our new gRPC interface.

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/telemetry.proto
```
