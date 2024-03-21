package piscine

func ThirdTimeIsACharm(s string) string {
	if s == "" || len(s) < 3 {
		return "\n"
	}
	result := ""
	for i := 2; i <= len(s)-1; i += 3 {
		result = result + string(s[i])
	}
	return result
}
