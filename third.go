package piscine

func Third(s string)string {
	result := ""
	for i := 2; i < len(s); i+=3 {
		result = result + string(s[i])
	}
	return result
}