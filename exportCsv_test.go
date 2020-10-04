package main

import (
	"os"
	"testing"
)

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Test_exportCsv(t *testing.T) {
	filePath := "./test.csv"
	fileInfo := FileInfo{
		"./abc.js",
		123,
		432,
		654,
	}

	fileInfos := []FileInfo {
		fileInfo,
	}

	exportCsv(fileInfos, filePath)

	if fileExists(filePath) == false {
		t.Error("Should have printed out file")
	}

	os.Remove(filePath)
}