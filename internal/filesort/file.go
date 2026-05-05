package filesort

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func SortFiles(groups map[string][]string, path string) error {
	var errs []error
	groupKeys := GetGroupKeys(groups)
	SortGroupKeysByDepth(&groupKeys)
	for _, key := range groupKeys {
		group := groups[key]
		if len(group) == 0 {
			continue
		}
		os.MkdirAll(filepath.Join(path, key), 0755)
		for _, pattern := range group {
			fileNameList, err := getFileListByPattern(pattern, path)
			if err != nil {
				errs = append(errs, fmt.Errorf("error getting filenames: %w", err))
			}
			for _, fileName := range fileNameList {
				moveFile(fileName, path, filepath.Join(path, key))
				if err != nil {
					errs = append(errs, fmt.Errorf("error moving file: %w", err))
				}
			}
		}
	}

	return errors.Join(errs...)
}

func getFileListByPattern(pattern string, path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var matchingFiles []string
	for _, file := range files {
		matched, err := filepath.Match(pattern, file.Name())
		if err != nil {
			return nil, err
		}
		if matched {
			matchingFiles = append(matchingFiles, file.Name())
		}
	}
	return matchingFiles, nil
}

func moveFile(fileName string, origin string, dest string) error {
	originPath := filepath.Join(origin, fileName)
	destPath := filepath.Join(dest, fileName)
	err := os.Rename(originPath, destPath)
	if err != nil {
		return err
	}
	return nil
}
