package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 34 || r > 126 {
			return false
		}
	}
	return true
}
