package main

import (
	"bufio"
	"index/suffixarray"
	"os"
	"regexp"
)

type FileInfo struct {
	Filepath   string `json:"filepath"`
	LineCount  int    `json:"lineCount"`
	CodeLines  int    `json:"codeLines"`
	Complexity int    `json:"complexity"`
}

var codeBranchIdentifierRegex = "(?P<ifStatements>if\\s?\\()| (?P<logicalOR>\\|\\|) | (?P<logicalAND>\\&\\&) | (?P<nullishCoalescing>\\?\\?) | (?P<switchCases>case+s?.*) | (?P<optionalChaining>\\?\\..*) | (?P<ternary> .+\\?.+\\:.+)"

func complexityCounter(filepath string) (fileInfo FileInfo) {
	r, _ := os.Open(filepath)
	conditionalRegex := regexp.MustCompile(codeBranchIdentifierRegex)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	complexityCount := 0
	lineCount := 0
	codeLines := 0
	for scanner.Scan() {
		lineCount += 1

		line := scanner.Text()

		if line != "" {
			codeLines += 1
		}

		lineAsBytes := suffixarray.New([]byte(line))
		matches := lineAsBytes.FindAllIndex(conditionalRegex, -1)

		complexityCount += len(matches)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return FileInfo{
		filepath,
		lineCount,
		codeLines,
		complexityCount,
	}
}
