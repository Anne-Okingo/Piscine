package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	for i := len(params) - 1; i >= 0; i-- {
		arguments := params[i]
		for _, value := range arguments {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
