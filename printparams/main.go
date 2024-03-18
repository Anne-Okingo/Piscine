package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	for _, arguments := range params {
		for _, value := range arguments {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
