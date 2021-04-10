package factory

import (
	"github.com/ip-rw/ransack/ent"
	"github.com/ip-rw/ransack/ent/proto/entpb"
	"github.com/ip-rw/ransack/pkg/proto"
	"google.golang.org/grpc"
	"log"
)

var (
	Addr   = "127.0.0.1:5000"
	Client *ent.Client
)

func connect(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	return conn
}

func GetHashClient(addr string) proto.HashServiceClient {
	return proto.NewHashServiceClient(connect(addr))
}
func GetMachineClient() proto.HashServiceClient {
	ent.NewClient()
	//return proto.NewHashServiceClient(connect(Addr))
}
