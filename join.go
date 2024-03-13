package piscine

func Join(strs []string, sep string) string {
	var result string
	for i, elem := range strs {
		result = result + elem
		if i != len(strs)-1 {
			result = result + sep
		}
	}
	return result
}
