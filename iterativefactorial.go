package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 10 {
		return 0
	} else if nb == 0 {
		return 1
	}
	result := 1
	for a := nb; a >= 1; a-- {
		result = result * a
	}
	return result
}
