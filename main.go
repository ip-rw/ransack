package main

import (
	"os"
	"github.com/ip-rw/ransack/pkg/walk"
)

func main() {
	walk.Walk([]string{"/home/none"}, func(dir string, info os.FileInfo, err error) error {

		return nil
	})
}
