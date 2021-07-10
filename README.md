# Poller Server

The purpose of this application is to poll time-series data from our [serialreader-server](https://github.com/bartmika/soilreader-servier) application and save it to the [tstorage-server](https://github.com/bartmika/tstorage-server) application. The interval of time is every one minute.

## Prerequisites

You must have the following installed before proceeding. If you are missing any one of these then you cannot begin.

* ``Go 1.16.3``

## Installation

1. Get our latest code.

    ```bash
    go get -u github.com/bartmika/poller-server
    ```

2. The purpose of this application is to poll time-series data from our [Mikapod Soil Reader](https://github.com/mikaponics/mikapod-soil-reader) application and save it to the [Mikapod Storage](https://github.com/mikaponics/mikapod-storage) application. The interval of time is every one minute.

## Prerequisites

You must have the following installed before proceeding. If you are missing any one of these then you cannot begin.

* ``Go 1.12.7``

## Installation
1. Please visit the [sparkfunweathershield-arduino](https://github.com/bartmika/sparkfunweathershield-arduino) repository and setup the external device and connect it to your development machine.

2. Please visit the [serialreader-server](https://github.com/bartmika/serialreader-server) repository and setup that application on your device.

3. Please visit the [tstorage-server](https://github.com/mikaponics/mikapod-storage) repository and setup that application on your device.

4. Get our latest code.

    ```bash
    go get -u github.com/bartmika/poller-server
    ```

5. Setup our environment variable. Please change to the address of our remote server. Please adjust the sensor ID values based what was assugned frin the Mikaponics web service. If the data belongs to a different tenant please adjust the value, else leave as is!

    ```
    # Sensor IDs
    export POLLER_SERVER_SERIAL_READER_SERVER_ADDRESS=127.0.0.1
    export POLLER_SERVER_SERIAL_READER_SERVER_PORT=50052
    export POLLER_SERVER_TSTORAGE_SERVER_ADDRESS=127.0.0.1
    export POLLER_SERVER_TSTORAGE_SERVER_PORT=50051
    ```

6. Run our application.

    ```
    go run main.go serve
    ```

7. If the server successfully starts you should see a message in your **termnal**:

    ```
    2021/07/10 15:40:36 Synching with local time...
    2021/07/10 15:41:00 Synchronized with local time.
    2021/07/10 15:41:00 Poller is now running.
    ```
