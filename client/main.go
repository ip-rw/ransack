package main

import (
	"context"
	"fmt"
	"github.com/ip-rw/ransack/pkg/core"
	"github.com/ip-rw/ransack/pkg/factory"
	"github.com/ip-rw/ransack/pkg/hash"
	"github.com/ip-rw/ransack/pkg/proto"
	"github.com/ip-rw/ransack/pkg/walk"
	"github.com/paulbellamy/ratecounter"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	sem    = make(chan int, 100)
	wg     = &sync.WaitGroup{}
	fsh    *hash.FsHash
	client = factory.GetHashClient("127.0.0.1:50051")
)

func lookupHash(hash uint64) (*proto.LookupResult, error) {
	hashes := proto.Hashes{
		Hashes: []uint64{hash},
	}
	return client.Lookup(context.Background(), &hashes)
}

var finished = false

func hashSubmitter() chan *proto.DirectoryHash {
	var dhc = make(chan *proto.DirectoryHash, 1000)
	var buf = &proto.DirectoryHashes{}

	submit := func(b *proto.DirectoryHashes) {
		res, err := client.SubmitHashes(context.Background(), b)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		for _, d := range b.DirectoryHashes {
			fmt.Println("\t", d.Name, d.Size)
		}
		//fmt.Println(res, b.DirectoryHashes)
		buf = &proto.DirectoryHashes{}
	}
	go func() {
		wg.Add(1)
		defer wg.Done()

		for !finished {
			select {
			case dh := <-dhc:
				buf.DirectoryHashes = append(buf.DirectoryHashes, dh)
				if len(buf.DirectoryHashes) > 50 {
					submit(buf)
				}
			case <-time.After(time.Second):
				if len(buf.DirectoryHashes) > 0 {
					submit(buf)
				}
			}
		}
		if len(buf.DirectoryHashes) > 0 {
			submit(buf)
		}
	}()
	return dhc
}

func main() {
	dhChan := hashSubmitter()

	var dirHashChan = make(chan string, 1)
	var hashrate = ratecounter.NewRateCounter(time.Second * 5)
	roots := []string{"/home/none"}
	var fsh = hash.NewFsHash()

	go func() {
		for fsh.Scanning {
			time.Sleep(3 * time.Second)
			fmt.Printf("%d d/s\n", hashrate.Rate()/5)
		}
	}()

	fsh.Scan(roots, func(path string, hash uint64, length int) {
		// if hash not found
		hashrate.Incr(1)
	})
	close(dirHashChan)
	for _, path := range roots {
		walk.Walk(path, func(dir string, info os.FileInfo, err error) error {
			fullpath := filepath.Join(dir, info.Name())
			if info.IsDir() {
				ha := fsh.LookupPath(fullpath)
				kfe := fsh.LookupHash(ha)

				if ha == 0 {
					return walk.SkipDir
				}
				found, err := lookupHash(ha)
				if err != nil {
					fmt.Println(err)
				}

				for _, dh := range found.FoundHashes {
					if dh.Hash == ha {
						//fmt.Printf("found %x (%s) skipping\n", dh.Hash, dh.Name)
						return walk.SkipDir
					}
				}

				dhChan <- &proto.DirectoryHash{
					Hash: ha,
					Name: info.Name(),
					Size: kfe[0].Size,
				}

				//fmt.Println("done", fullpath)
			} else {
				core.SearchFile(fullpath, info)
			}
			return nil
		})
	}
	finished = true
	wg.Wait()
	//fmt.Println(fsh.SaveGob("fsh.bin"))
}
