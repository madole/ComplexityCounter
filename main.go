package main

import (
	"flag"
)

var extensionsRegex string

type FilesInfo []FileInfo

func initFlags() (string, string, string, string) {
	extensionsRegex := flag.String("extensionsRegex", "\\.(js|jsx|ts|jsx)$", "A regex for specifying file extensions")
	directory := flag.String("directory", "./", "The directory to search")
	exclusionsRegex := flag.String("exclusionsRegex", "(__tests__|test\\.(js|ts|jsx|tsx))|__snapshots__", "Any other files you dont want to include")
	outputType := flag.String("outputType", "csv", "Choose between csv or json")
	flag.Parse()
	return *extensionsRegex, *directory, *exclusionsRegex, *outputType
}

func main() {
	extensionsRegex, directory, exclusionsRegex, outputType := initFlags()

	filePaths, err := getFilePathsForDirectory(directory, extensionsRegex, exclusionsRegex)

	if err != nil {
		panic(err)
	}

	var files FilesInfo

	for _, file := range filePaths {
		fileInfo := complexityCounter(file)

		if err != nil {
			panic(err)
		}
		files = append(files, fileInfo)
	}

	if outputType == "csv" {
		exportCsv(files, "./output.csv")
	} else if outputType == "json" {
		exportJson(files, "./output.json")
	}
}
