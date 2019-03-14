package lineshift

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// arguments that provides access to command-line arguments.
func main() {
	arg1 := os.Args[1]
	MacOrWindowsLineShift(arg1)

}

// MacOrWindowsLineShift  opens a file and checking the OS.
func MacOrWindowsLineShift(filename string) {
	file1, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	data1, err := ioutil.ReadAll(file1)
	if err != nil {
		log.Fatal(err)
	}

	match1, err := regexp.Match(`\r\n`, data1)
	if err != nil {
		log.Fatal(err)
	}
	text1Platform := "Windows"
	if !match1 {
		text1Platform = "Mac"
	}

	fmt.Printf("The textfile is using %s line breaks\n", text1Platform)

}
