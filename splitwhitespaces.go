package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	var currentWord string

	for _, char := range s {
		switch char {
		case ' ', '\t', '\n':
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		default:
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	return words
}
