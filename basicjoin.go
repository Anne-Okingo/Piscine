package piscine

func BasicJoin(elems []string) string {
	var j string
	for _, char := range elems {
		j += char
	}
	return j
}
