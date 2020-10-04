package main

func printFileInfo(fileInfo FileInfo) {
	println(fileInfo.filepath)
	print("line count: ")
	println(fileInfo.lineCount)
	print("code lines: ")
	println(fileInfo.codeLines)
	print("complexity: ")
	println(fileInfo.complexity)
}
