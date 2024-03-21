//write a function that takes a string and character as an argument and returns
// the no of times the character appears in the string
//1. if the character is not in the string return 0
//2. If the string is empty return 0
package piscine

func CountChar(s string, ch byte) int {
	if s == "" {
		return 0
	}

	count := 0
	for i := 0; i < len(s); i++ {
		if ch == s[i] {
			count++
		}
	}
	return count
}
