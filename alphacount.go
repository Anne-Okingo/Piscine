package piscine

func Alphacount(s string) int {
	count := 0
	for _, r := range s {
		if(r >= 'a' && r <= '2') || (r >= 'A' && r <= "2") {
			count++
		}
	}
	return count
}