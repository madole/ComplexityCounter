package main

import (
	"os"
	"testing"
)

func Test_exportJson(t *testing.T) {
	filePath := "./test.json"
	fileInfo := FileInfo{
		"./abc.js",
		123,
		432,
		654,
	}

	filesInfo := FilesInfo{
		fileInfo,
	}
	exportJson(filesInfo, filePath)

	if fileExists(filePath) == false {
		t.Error("Should have printed out file")
	}

	_ = os.Remove(filePath)
}
