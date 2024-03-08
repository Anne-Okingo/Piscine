package piscine

func ToUpper(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= '2' {
			b[i] = c - ('a' - 'A')
		}
	}
	return string(b)
}
