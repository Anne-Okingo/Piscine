package piscine

func CountChar(s string, ch byte) int {
	if s == "" {
		return 0
	}

	count := 0
	for i := 0; i < len(s); i++ {
		if ch == s[i] {
			count++
		}
	}
	return count
}
