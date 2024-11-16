package main

import (
	"log"
	"net"
	"net/rpc"

"github.com/aman-saraiya/remote-procedure-calls/rpc-go-standard/server/userservice"
)

func main() {
	
	// Declaring and Initializing the UserService
	userSvc := userservice.NewUserService()

	// Registering the UserService service to the RPC server
	err := rpc.Register(userSvc)
	if err != nil {
		log.Fatalf("Error registering UserService: %v", err)
	}

	// Starting a TCP listener at port 2024
	listener, err := net.Listen("tcp", "localhost:2024")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	log.Println("Server is listening on port 2024")
	defer listener.Close()

	for {
		// Accept incoming connection requests from clients
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Serve RPC over the accepted client connection 
		go rpc.ServeConn(conn)
	}
}
