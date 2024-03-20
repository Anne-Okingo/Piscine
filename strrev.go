package piscine

func StrRev(s string) string {
	p := []rune(s)
	   t := []rune{}
	   length := len(p)
	   for i := length - 1; i >= 0; i-- {
		    t = append(t, p[i])
	}
	return string(t)
}
