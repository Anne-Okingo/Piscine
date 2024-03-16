package piscine

func StrRev(s string) string {
	runes := []rune(s)
	reversed := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}
	return string(reversed)
}
