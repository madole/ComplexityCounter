package main

import "testing"

func Test_ComplexityCounter(t *testing.T) {
	filePath := "./TestFile.js"
	got := complexityCounter(filePath)

	t.Run("Testing line count", func(st *testing.T) {
		expectedLines := 50
		if got.LineCount != expectedLines {
			t.Errorf("Should have %d lines", expectedLines)
		}
	})

	t.Run("Testing code lines count", func(st *testing.T) {
		expectedCodeLines := 48
		if got.CodeLines != expectedCodeLines {
			t.Errorf("Should have %d", expectedCodeLines)
		}
	})

	t.Run("Testing file path", func(st *testing.T) {
		if got.Filepath != filePath {
			t.Errorf("Should have been %s", filePath)
		}
	})

	t.Run("Testing complexity", func(st *testing.T) {
		expectedComplexity := 7
		if got.Complexity != expectedComplexity {
			t.Errorf("Should have been %d", expectedComplexity)
		}
	})
}
