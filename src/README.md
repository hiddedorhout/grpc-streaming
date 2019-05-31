# gRPC Server Side streaming testing

The goal of this project is to test server sided streaming using the gRPC server, a javascript client application with `grpc-web` and the `envoy proxy`. Works on Linux only (for now).

## Prerequisites

* Having the `protoc` compiler locally installed with the `golang` and `protoc-gen-grpc-web` plugin.
* Docker
* Go (Run this project in GOPATH)
* npm

## Run

1. From this directory: `make srcGen`
2. Run Envoy proxy (linux only): `make run-proxy`
3. Run server: `make run-server`
4. Run client:
    1. Go to `js/client`
    2. Install: `npm install`
    3. Run: `npm start`

The screen should display a state change based on the stream from the server.