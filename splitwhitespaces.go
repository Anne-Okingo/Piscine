package piscine

func SplitWhiteSpaces(s string) []string {
	returnedArray := []string{}
	startPoint := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			returnedArray = append(returnedArray, s[startPoint:i])
			startPoint = i + 1
		}
	}
	returnedArray = append(returnedArray, s[startPoint:])
	return returnedArray
}
