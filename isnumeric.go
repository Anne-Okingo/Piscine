package piscine

func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			continue
		} else if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			return false
		} else {
			return false
		}
	}
	return true
}
