package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 00; a <= 98; a++ {
		b := 00 + 1; b <= 99; b++ {
			if a != 98 && a != 99 {
			z01.PrintRune(a)
			z01.PrintRune(b)
			}
			z01.PrintRune('\n')
		}
	}
}