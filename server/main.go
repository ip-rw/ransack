package main

import (
	"context"
	"fmt"
	"github.com/ip-rw/ransack/pkg/hash"
	"github.com/ip-rw/ransack/pkg/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	proto.UnimplementedHashServiceServer
}

var fsh *hash.FsHash

func (s *server) Lookup(context context.Context, hashes *proto.Hashes) (*proto.LookupResult, error) {
	log.Printf("Received: %v", hashes.GetHashes())
	res := &proto.LookupResult{
		RequestedHashes: hashes.GetHashes(),
		FoundHashes:     []*proto.DirectoryHash{},
	}
	for _, dh := range hashes.GetHashes() {
		if kfe := fsh.LookupHash(dh); kfe != nil {
			res.FoundHashes = append(res.FoundHashes, &proto.DirectoryHash{
				Hash: dh,
				Name: kfe[0].Path,
				Size: kfe[0].Size,
			})
		}
	}
	return res, nil
}

func (s *server) SubmitHashes(context context.Context, hashes *proto.DirectoryHashes) (*proto.Result, error) {
	fsh.Lock()
	defer fsh.Unlock()
	for _, dh := range hashes.DirectoryHashes {
		if _, ok := fsh.Hmap[dh.Hash]; !ok {
			fsh.Hmap[dh.Hash] = []hash.DirectoryHash{}
		}
		fsh.Hmap[dh.Hash] = append(fsh.Hmap[dh.Hash], hash.DirectoryHash{
			Path: dh.Name,
			Size: dh.Size,
			Hash: dh.Hash,
		})
	}
	return &proto.Result{
		Success:  true,
		Affected: int64(len(hashes.DirectoryHashes)),
	}, nil
}

func main() {
	if fsh, _ = hash.Load("fsh.bin"); fsh == nil {
		fsh = hash.NewFsHash()
	}
	go func() {
		for {
			time.Sleep(30 * time.Second)

			fsh.SaveGob("fsh.bin")
			fmt.Println("saved fsh.bin")
		}
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterHashServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
