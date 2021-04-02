package main

import (
	"context"
	"github.com/ip-rw/ransack/pkg/hash"
	"log"
	"net"

	pb "github.com/ip-rw/ransack/pkg/proto/hash"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

var fsh *hash.FsHash

func init() {
	if fsh, _ := hash.Load("fsh.bin"); fsh == nil {
		fsh = hash.NewFsHash([]string{})
	}
}
func (s *server) SubmitHashes(ctx context.Context, in *pb.DirectoryHashes) (*pb.Result, error) {
	log.Printf("Received: %v", in.DirectoryHashes)
	for _, dh := range in.DirectoryHashes {
		fsh
	}
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
