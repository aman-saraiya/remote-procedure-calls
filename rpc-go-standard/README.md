# RPC Standard Go Implementation

Link to Medium Article: [**gRPC is not the only RPC framework: Exploring Remote Procedure Calls**](#link-to-article-1)

This directory contains a simple implementation of standard RPC in Go. It demonstrates basic concepts of RPC, including how to create a server and a client that communicate using Go's built-in `net/rpc` package.

## Directory Structure

- **`client/`**: Contains the client code for making RPC calls to the server.
- **`server/`**: Contains the server code that listens for and responds to RPC calls.
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

2. **Run the Client**:
   - In a new terminal, navigate to the `client` directory:
     ```bash
     cd client
     ```
   - Run the pre-built client binary:
     ```bash
     ./client
     ```

The server should now be running, and the client can make RPC calls to it.

## Setup Requirements to Build

- **Go 1.23.3** or higher is required to build the source code. Make sure Go is installed on your system before attempting to build the code.

## Building the Code

If you want to modify the source code or build the code for a different platform, follow these steps:

1. **Make Updates**: Modify the source code locally in either the `client` or `server` directories.

2. **Build the Binaries**:
   - After making your changes, navigate to each directory (`client` or `server`).
   - Run the following command to build the binaries:
     ```bash
     go build .
     ```

This will generate new binary files that you can run directly.
