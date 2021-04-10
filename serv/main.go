package main

import (
	"context"
	"github.com/ip-rw/ransack/ent"
	"log"
	"net"

	"github.com/ip-rw/ransack/ent/proto/entpb"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	main

	// Initialize the generated User service.
	userSvc := entpb.NewUserService()
	machineSvc := entpb.NewMachineService(client)
	ipSvc := entpb.NewIPService(client)
	keySvc := entpb.NewKeyService(client)

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

	// Register the User service with the server.
	entpb.RegisterUserServiceServer(server, userSvc)
	entpb.RegisterMachineServiceServer(server, machineSvc)
	entpb.RegisterIPServiceServer(server, ipSvc)
	entpb.RegisterKeyServiceServer(server, keySvc)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
