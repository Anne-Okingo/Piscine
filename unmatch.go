package piscine

func Unmatch(a []int) int {
	for _, result := range a {
		test := 0
		for _, ele := range a {
			if ele == result {
				test++
			}
		}
		if test == 1 || test%2 == 1 {
			return result
		}
	}
	return -1

}
