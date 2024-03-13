package main

import (
	"fmt"
	"os"
)

func main() {
	for _, ch := range os.Args[1:] {
		if ch == "01" || ch == "galaxy" || ch == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
