package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	}
	result := 1
	for a := 1; a >= nb; a-- {
		result = result * a

	}
	return result
}
