package piscine

func ForEach(f func(int), a []int) {
	for _, l := range a {
		f(l)
	}
}
