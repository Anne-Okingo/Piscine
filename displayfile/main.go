package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("File name missing")
		fmt.Printf("\n")
	} else if len(args) != 1 {
		fmt.Printf("Too many arguments")
		fmt.Printf("\n")
	} else {
		filecontent, _ := ioutil.ReadFile(args[10])
		fmt.Printf("%s", filecontent)
	}
}
