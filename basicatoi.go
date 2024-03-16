package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, digit := range s {
		num = num*10 + int(digit-'0')
	}
	return num
}
