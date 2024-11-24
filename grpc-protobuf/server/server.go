package main

import (
	"log"
	"net"

	gen_statssvc "github.com/aman-saraiya/remote-procedure-calls/grpc-protobuf/client/generated/statsservice"
	gen_usersvc "github.com/aman-saraiya/remote-procedure-calls/grpc-protobuf/client/generated/userservice"
	"github.com/aman-saraiya/remote-procedure-calls/grpc-protobuf/server/statsservice"
	"github.com/aman-saraiya/remote-procedure-calls/grpc-protobuf/server/userservice"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	userSvcServer := userservice.NewUserServiceServer()
	statsSvcServer := statsservice.NewStatsServiceServer()

	gen_usersvc.RegisterUserServiceServer(grpcServer, userSvcServer)
	gen_statssvc.RegisterStatsServiceServer(grpcServer, statsSvcServer)

	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
