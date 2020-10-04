package main

/**
* DEBUG FUNCTION
**/
func printFileInfo(fileInfo FileInfo) {
	println(fileInfo.Filepath)
	print("line count: ")
	println(fileInfo.LineCount)
	print("code lines: ")
	println(fileInfo.CodeLines)
	print("complexity: ")
	println(fileInfo.Complexity)
}
