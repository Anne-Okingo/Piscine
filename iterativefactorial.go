package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		a := nb - 1
		for b := a; b >= 1; b-- {
			nb *= b
		}
	}
	return nb
}
