package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func exportJson(filesInfo FilesInfo, outputPath string) {
	jsonOutput, err := json.MarshalIndent(filesInfo, "", " ")
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile(outputPath, jsonOutput, os.ModePerm)
}