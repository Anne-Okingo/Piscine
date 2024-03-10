package piscine

func ConcatParams(arg []string) string {
	result := arg[0]
	for _, s := range arg[1:] {
		result = result + "\n" + s
	}
	return result
}
