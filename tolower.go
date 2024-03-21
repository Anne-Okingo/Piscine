package piscine

func ToLower(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			result = result + string(ch-'A'+'a')
		} else {
			result = result + string(ch)
		}
	}
	return result
}
