package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func getFileInfo(filnavn string) {
	file, err := os.Lstat(filnavn)

	if err != nil {
		log.Fatal(err)
	}
	mode := file.Mode()
	//converts to float64 for accurracy
	size := float64(file.Size())
	fmt.Println("Information about:", file.Name())
	dir, _ := os.Getwd()
	fmt.Println("path:", dir)
	fmt.Println("Unix Permissions Bits:", mode.Perm())
	fmt.Printf("Directory: %t\n", mode.IsDir())
	fmt.Printf("Regular: %t\n", mode.IsRegular())
	fmt.Printf("Size: %v bytes, %v kibibytes, %v mibibytes and %v gibibytes\n",
		size, size/1024, size/(1024*1024), size/(1024*1024*1024))
	fmt.Printf("Symbolic link: %t\n", mode&os.ModeSymlink != 0)
	fmt.Printf("Append only: %t\n", mode&os.ModeAppend != 0)
	fmt.Printf("Device file: %t\n", mode&os.ModeDevice != 0)
	fmt.Printf("Unix char device: %t\n", os.ModeDevice&os.ModeCharDevice != 0)

}
func main() {

	var filnavn string

	flag.StringVar(&filnavn, "f", "", "Name of file to be inspected")

	flag.Parse()

	getFileInfo(filnavn)

}
