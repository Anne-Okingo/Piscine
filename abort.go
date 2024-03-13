package piscine

func Abort(a, b, c, d, e int) int {
	k := []int{a, b, c, d, e}
	for i := 0; i < len(k); i++ {
		for j := i + 1; j < len(k); j++ {
			if k[i] > k[j] {
				k[i], k[j] = k[j], k[i]
			}
		}
	}
	return k[2]
}
