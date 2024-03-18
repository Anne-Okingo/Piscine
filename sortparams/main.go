package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	for i := 0; i < len(params)-1; i++ {
		for j := i + 1; j < len(params); j++ {
			if params[j] < params[i] {
				params[i], params[j] = params[j], params[i]
			}
		}
	}
	for k := 0; k < len(params); k++ {
		for _, value := range params[k] {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
