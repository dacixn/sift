package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// usage should be
// sift .

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
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		dir = filepath.Join(wd, args[1])
	} else {
		err = errors.New("Too many arguments")
		log.Fatal(err)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fileMap := groupFilesByExtension(files)
	sortFiles(fileMap, dir)
}

func groupFilesByExtension(files []os.DirEntry) map[string][]string {
	fileMap := make(map[string][]string)
	for _, file := range files {
		fileType, ok := getFileType(file)
		if !ok {
			fmt.Printf("Skipping %s\n", file.Name())
			continue
		}
		fileMap[fileType] = append(fileMap[fileType], file.Name())
	}
	return fileMap
}

func getFileType(file os.DirEntry) (string, bool) {
	fileType := filepath.Ext(file.Name())
	if fileType == "" || file.IsDir() {
		return "", false
	}
	return fileType, true
}

func moveFileToExtensionFolder(fileName, sourceDir, ext string) error {
	os.MkdirAll(filepath.Join(sourceDir, ext), 0755)
	oldPath := filepath.Join(sourceDir, fileName)
	newPath := filepath.Join(sourceDir, ext, fileName)
	return os.Rename(oldPath, newPath)
}

func sortFiles(fileMap map[string][]string, sourceDir string) error {
	for ext, group := range fileMap {
		for _, file := range group {
			err := moveFileToExtensionFolder(file, sourceDir, ext)
			if err != nil {
				fmt.Printf("Error moving file: %v", err)
			}
		}
	}
	return nil
}
