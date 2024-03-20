package piscine

func ToUpper(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result = result + string(ch-'a'+'A')
		} else {
			result = result + string(ch)
		}
	}
	return result
}
...