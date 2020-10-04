package main

import "testing"

func Test_GetFilePathsForDirectory(t *testing.T) {
	t.Run("Returns the correct number of matched file paths", func(st *testing.T) {
		got, _ := getFilePathsForDirectory("./", "\\.(js|jsx|ts|jsx)", "(__tests__|test\\.(js|ts|jsx|tsx))|__snapshots__")
		expectedLength := 1
		if len(got) != expectedLength {
			t.Errorf("Expected to get %d record returned", expectedLength)
		}
	})

	t.Run("Deals with empty string", func(st *testing.T) {
		_, err := getFilePathsForDirectory("", "\\.(js|jsx|ts|jsx)", "(__tests__|test\\.(js|ts|jsx|tsx))|__snapshots__")
		if err == nil {
			t.Error("Should have thrown an error")
		}
	})
}
