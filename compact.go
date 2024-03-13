package piscine

func Compact(ptr *[]string) int {
	new := *ptr
	look := 0
	for _, ch := range *ptr {
		if ch != "" {
			new[look] = ch
			look++
		}
	}
	*ptr = new[0:look]
	return look
}
