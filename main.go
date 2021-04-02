package main

import (
	"fmt"
	"github.com/ip-rw/ransack/pkg/hash"
	"github.com/ip-rw/ransack/pkg/walk"
	"github.com/paulbellamy/ratecounter"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	sem = make(chan int, 100)
	wg  = &sync.WaitGroup{}
	fsh *hash.FsHash
)

func main() {
	var dirHashChan = make(chan string, 1)
	var hashrate = ratecounter.NewRateCounter(time.Second * 5)
	roots := []string{"/home/none"}
	var fsh *hash.FsHash
	var err error
	if fsh, err = hash.Load("fsh.bin"); err != nil {
		fmt.Println(err)
		fsh = hash.NewFsHash(roots)
	}
	go func() {
		for fsh.Scanning {
			time.Sleep(3 * time.Second)
			fmt.Printf("%d d/s\n", hashrate.Rate()/5)
		}
	}()

	fsh.Scan(func(path string, hash uint64, length int) {
		// if hash not found
		hashrate.Incr(1)
	})
	close(dirHashChan)
	for _, path := range roots {
		walk.Walk(path, func(dir string, info os.FileInfo, err error) error {
			fullpath := filepath.Join(dir, info.Name())
			if info.IsDir() {
				ha := fsh.LookupPath(fullpath)
				if ha == 0 {
					return walk.SkipDir
				}
				found := fsh.LookupHash(ha)
				if len(found) > 1 {
					if dir != found[0].Path {
						if found[0].Size > 0 {
							fmt.Printf("Skipping %s, one of %d duplicates of %s (%d bytes)\n", fullpath, len(found), found[0].Path, found[0].Size)
						}
						return walk.SkipDir
					}
				}
			}
			return nil
		})
	}
	wg.Wait()
	fmt.Println(fsh.SaveGob("fsh.bin"))
}
