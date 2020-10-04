package main

import (
	"flag"
)

type FilesInfo []FileInfo

func initFlags() (string, string, string, string, bool) {
	extensionsRegex := flag.String("extensionsRegex", "\\.(js|jsx|ts|jsx)$", "A regex for specifying file extensions")
	directory := flag.String("directory", "./", "The directory to search")
	exclusionsRegex := flag.String("exclusionsRegex", "(__snapshots__|__tests__|test\\\\.(js|ts|jsx|tsx))", "Any other files you dont want to include")
	outputType := flag.String("outputType", "csv", "Choose between csv or json")
	debug := flag.Bool("debug", false, "Turn on to log file info")
	flag.Parse()
	return *extensionsRegex, *directory, *exclusionsRegex, *outputType, *debug
}

func main() {
	extensionsRegex, directory, exclusionsRegex, outputType, debug := initFlags()

	filePaths, err := getFilePathsForDirectory(directory, extensionsRegex, exclusionsRegex)

	if err != nil {
		panic(err)
	}

	var files FilesInfo

	for _, file := range filePaths {
		fileInfo := complexityCounter(file)
		files = append(files, fileInfo)
		if debug == true {
			printFileInfo(fileInfo)
		}
	}

	if outputType == "csv" {
		exportCsv(files, "./output.csv")
	} else if outputType == "json" {
		exportJson(files, "./output.json")
	} else {
		panic("Output type must be csv or json")
	}
}
