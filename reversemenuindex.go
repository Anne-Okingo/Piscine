package piscine

func ReverseMenuIndex(menu []string) []string {
	argt := make([]string, len(menu))
	for i := len(menu) - 1; i >= 0; i-- {
		argt[len(menu)-1-i] = menu[i]
	}
	return argt
}
