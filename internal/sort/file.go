package filesort

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// func GroupFilesByExtension(files []os.DirEntry) map[string][]string {
// 	fileMap := make(map[string][]string)
// 	for _, file := range files {
// 		fileType, ok := GetFileType(file)
// 		if !ok {
// 			continue
// 		}
// 		fileMap[fileType] = append(fileMap[fileType], file.Name())
// 	}
// 	return fileMap
// }

// func GetFileType(file os.DirEntry) (string, bool) {
// 	fileType := filepath.Ext(file.Name())
// 	if fileType == "" || file.IsDir() {
// 		return "", false
// 	}
// 	return fileType, true
// }

func MoveFileToExtensionFolder(fileName, sourceDir, ext string) error {
	oldPath := filepath.Join(sourceDir, fileName)
	newPath := filepath.Join(sourceDir, ext, fileName)
	return os.Rename(oldPath, newPath)
}

func SortFiles(fileMap map[string][]string, sourceDir string) error {
	var errs []error
	for ext, group := range fileMap {
		err := os.MkdirAll(filepath.Join(sourceDir, ext), 0755)
		if err != nil {
			errs = append(errs, fmt.Errorf("could not create directory for %s: %w", ext, err))
			continue
		}
		for _, file := range group {
			err := MoveFileToExtensionFolder(file, sourceDir, ext)
			if err != nil {
				errs = append(errs, fmt.Errorf("could not move %s: %w", file, err))
			}
		}
	}
	return errors.Join(errs...)
}
