package piscine

func ToLower(s string) string {
	var result string
	for _, char := range s {
		if char >= 'A' && char <= 'z' {
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}
	return result
}
