package piscine

func JumpOver(str string) string {
	stri := ""
	if len(str) >= 3 {
		for i, ch := range str {
			if i%3 == 2 {
				stri = stri + string(ch)
			}
		}
	}
	return stri + "\n"
}
