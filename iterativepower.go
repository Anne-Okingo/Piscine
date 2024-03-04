package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	result := 1
	for a := 1; a <= power; a++ {
		result = result * nb
	}
	return result
}
