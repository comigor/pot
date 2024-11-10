## Pot

pot is a code generation tool for Go that simplifies HTTP handler generation using Protobuf files, much like protoc-gen-go-grpc, but specifically for HTTP. Unlike setups that require wrapping a gRPC handler with an HTTP gateway, pot directly generates HTTP handlers from Protobuf definitions. This provides an efficient, direct path for HTTP handling without gRPC, reducing complexity and dependency on gRPC gateways.

This project is inspired by go-kratos, a Go microservices framework. Special thanks to the go-kratos team for their foundational ideas that helped shape pot.

### Features

HTTP handler generation: Generates HTTP handlers directly from Protobuf service definitions.
gRPC independence: No need to wrap handlers with gRPC gateways; useful for services focused purely on HTTP.
Consistent API contract: Reuses Protobuf contracts, providing consistency across handler generation tools.
Motivation
Traditional setups often involve creating a gRPC handler and then wrapping it with a gateway like gRPC Gateway to serve both gRPC and HTTP clients. However, this approach can add unnecessary layers and dependencies for projects that only need HTTP services. pot fills this gap by generating HTTP handlers directly, streamlining HTTP-based development without compromising on API contract consistency.

### Installation

To install pot, you can simply use go install:

```sh
go install github.com/afikrim/pot@latest
```

### Usage

To generate HTTP handlers from your Protobuf files, run the following command:

```sh
protoc --plugin=protoc-gen-pot --pot_out=. --proto_path=PATH_TO_PROTO_FILE your_service.proto
```

Replace PATH_TO_PROTO_FILE with the path to your .proto files, and your_service.proto with the specific Protobuf file to generate handlers for.

#### Example

Given a Protobuf file example.proto, use pot as follows:

```sh
protoc --plugin=protoc-gen-pot --pot_out=. --proto_path=proto example.proto
```

This generates an HTTP handler based on the service definition in example.proto.

### Credits

pot was inspired by the design and ideas of the [go-kratos](https://github.com/go-kratos/kratos/tree/main/cmd/protoc-gen-go-http) project. We extend our gratitude to the [go-kratos](https://github.com/go-kratos/kratos/tree/main/cmd/protoc-gen-go-http) team for their innovation in microservices frameworks and their contributions to the Go community.

### License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/afikrim/pot/blob/main/LICENSE) file for more details.
