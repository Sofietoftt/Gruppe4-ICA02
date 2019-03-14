package main

import (
	"fmt"

	"github.com/tuvace/fileutils"
)

func main() {
	file1 := fileutils.FileToByteslice("files/text1.txt")
	file2 := fileutils.FileToByteslice("files/text2.txt")

	fmt.Println(file1)
	fmt.Println(file2)
}
