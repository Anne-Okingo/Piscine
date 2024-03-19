package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for r := 99; r >= 1; r-- {
		for s := r - 1; s >= 0; s-- {
			z01.PrintRune(rune(r/10) + '0')
			z01.PrintRune(rune(r%10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(s/10) + '0')
			z01.PrintRune(rune(s%10) + '0')
			if r != 1 || s != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}

	}
}
