package main

import (
	"os"
	"strconv"
	"strings"
)

var headers = []string{"file path", "line count", "code lines", "complexity"}

func exportCsv(filesInfo []FileInfo, outputFilepath string) {
	csvOutput := strings.Join(headers, ",")
	for _, fileInfo := range filesInfo {
		csvLine := "\n"
		csvLineData := []string{fileInfo.filepath, strconv.Itoa(fileInfo.lineCount), strconv.Itoa(fileInfo.codeLines), strconv.Itoa(fileInfo.complexity)}
		csvLine += strings.Join(csvLineData, ",")
		csvOutput += csvLine
	}
	outputFile, err := os.Create(outputFilepath)
	if err != nil {
		panic(err)
	}

	_, outputError := outputFile.WriteString(csvOutput)

	if outputError != nil {
		panic(err)
	}
}