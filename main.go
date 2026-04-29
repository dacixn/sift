package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// usage should be
// sift .

func main() {
	args := os.Args

	if len(args) == 1 {

	}
	fmt.Println(os.Getwd())
	dir := os.Getwd()
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fileMap := groupFilesByExtension(files)
	sortFiles(fileMap, dir)

	jsonMap, err := json.Marshal(fileMap)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("fileMap.json", jsonMap, 0755)

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
