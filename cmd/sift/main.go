package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/dacixn/sift/internal/sort"
)

func main() {
	var dir string
	var err error

	args := os.Args
	if len(args) == 1 {
		dir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	} else if len(args) == 2 {
		dir = filepath.Join(args[1])
	} else {
		err = errors.New("Too many arguments")
		log.Fatal(err)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fileMap := sort.GroupFilesByExtension(files)
	err = sort.SortFiles(fileMap, dir)
	if err != nil {
		log.Fatal(err)
	}
}
