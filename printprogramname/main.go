package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	progName := os.Args[0][2:]
	for _, ch := range progName {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
