package main

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"regexp"
)

func getFilePathsForDirectory(directory string, extensionsRegex string, exclusionsRegex string) ([]string, error) {
	if directory == "" {
		return nil, errors.New("Directory cannot be an empty string")
	}
	var filePaths []string

	extensionRegex, _ := regexp.Compile(extensionsRegex)
	excludeRegex, _ := regexp.Compile(exclusionsRegex)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && extensionRegex.MatchString(path) && !excludeRegex.MatchString(path) {
			filePaths = append(filePaths, path)
		}

		return nil
	})
	return filePaths, err
}
