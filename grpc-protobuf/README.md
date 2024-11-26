# gRPC with Protobuf Go Implementation

Link to Medium Article: [**A Deep Dive into gRPC: Service Definitions, Protobufs, and Hands-On Implementation**](https://medium.com/@aman-saraiya/a-deep-dive-into-grpc-service-definitions-protobufs-and-hands-on-implementation-777946d733ca)

This directory contains an implementation of gRPC services using Protocol Buffers (Protobuf) for service definitions, message serialization, and communication in Go. It demonstrates how to set up gRPC with Protobuf and includes both pre-built binaries and instructions to build the code from source.

## Directory Structure

- **`client/`**: Contains the client code that communicates with the server using gRPC.
- **`server/`**: Contains the server code that implements the gRPC services and handles requests.
- **`client/protos/`**: Contains the `.proto` files defining the gRPC services and message types.
- **`README.md`**: This file.

## Running the Code

This implementation includes pre-built binary files for Linux, which can be directly executed. To run the code:

1. **Start the Server**:

   - Navigate to the `server` directory:
     ```bash
     cd server
     ```
   - Run the pre-built server binary:
     ```bash
     ./server
     ```
     The server will start and begin listening for incoming gRPC requests.

2. **Run the Client**:
   - In a new terminal, navigate to the `client` directory:
     ```bash
     cd client
     ```
   - Run the pre-built client binary:
     ```bash
     ./client
     ```
     The client will initiate gRPC requests to the server.

### Setup Requirements to Build

- **Go 1.23.3** or higher is required to build the source code. Make sure Go is installed on your system before attempting to build the code.

- Before you can compile the Protobuf files and generate the necessary Go and gRPC code, you need to install the following tools:

  **Install the Protocol Buffers compiler plugin for Go**:

  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ```

  **Install the Protocol Buffers compiler plugin for gRPC Go**:

  ```bash
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

## Building the Code

> Update the proto files

Compiling the Protobuf Files

Once the setup requirements are complete, you can generate the Go and gRPC code from the .proto files. To do this:

Navigate to the client/protos directory:

```bash
cd client/protos
```

Run the following commands to compile the Protobuf files and generate the Go code for each service:

For statsservice.proto:

```bash
protoc --go_out=../ --go-grpc_out=../ statsservice.proto
```

For userservice.proto:

```bash
protoc --go_out=../ --go-grpc_out=../ userservice.proto
```

These commands will generate the corresponding .pb.go and \_grpc.pb.go files, which are necessary for the client and server to communicate over gRPC.

> Update the client and the server code

Build the Binaries:

After making your changes, navigate to each directory (client or server).
Run the following command to build the binaries:

```bash
go build .
```

This will generate new binary files that can be executed directly.
